const request = require('request');
const jwt = require('jsonwebtoken');

class TokenHelper {
  // This token generation is relevant to only Multitenant SaaS mode deployment.
  static async generateTokenHttp(clientId, clientSecret, url) {
    if (!clientId || !clientSecret) {
      throw new Error("Client Id or Secret is invalid");
    }

    const data = `grant_type=client_credentials&client_id=${clientId}&client_secret=${clientSecret}`;
    const options = {
      url,
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: data
    };

    return new Promise((resolve, reject) => {
      request(options, (error, response, body) => {
        if (error) {
          reject(error);
        } else if (response.statusCode >= 400) {
          reject(new Error(`Error in fetching token. ${body}`));
        } else {
          const accessToken = JSON.parse(body).access_token;
          resolve(accessToken);
        }
      });
    });
  }

  static isTokenExpired(accessToken) {
    const decoded = jwt.decode(accessToken);

    if (!decoded || !decoded.exp) {
      return true;
    }

    const expTimestamp = decoded.exp;
    const expDateTimeOffset = new Date(expTimestamp * 1000);
    return Date.now() >= expDateTimeOffset;
  }

  static async fetchToken(currentToken, clientId, clientSecret, url) {
    if (currentToken && !TokenHelper.isTokenExpired(currentToken)) {
      return currentToken;
    } else {
      return await TokenHelper.generateTokenHttp(clientId, clientSecret, url);
    }
  }
}

module.exports = TokenHelper;