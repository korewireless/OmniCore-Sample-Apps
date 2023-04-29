using System.Text.Json;
using System.IdentityModel.Tokens.Jwt;

namespace Example
{
    public static class TokenHelper
    {
        // This token generation is relevant to only Multitenant SaaS mode deployment.
        private static async Task<string> GenerateTokenHttp(string clientId, string clientSecret, string url)
        {
            if (string.IsNullOrWhiteSpace(clientId) || string.IsNullOrWhiteSpace(clientSecret))
            {
                throw new Exception("Client Id or Secret is invalid");
            }

            var data = $"grant_type=client_credentials&client_id={clientId}&client_secret={clientSecret}";
            var content = new StringContent(data, System.Text.Encoding.UTF8, "application/x-www-form-urlencoded");
            using var client = new HttpClient();
            var response = await client.PostAsync(url, content);

            if (response.IsSuccessStatusCode)
            {
                var responseString = await response.Content.ReadAsStringAsync();
                var accessToken = JsonSerializer.Deserialize<AccessToken>(responseString).access_token;
                return accessToken;
            }
            else
            {
                var error = await response.Content.ReadAsStringAsync();
                Console.WriteLine ("CleintId:" + clientId );
                throw new Exception($"Error in fetching token. {error}");
            }
        }

        private static bool IsTokenExpired(string accessToken)
        {
            var handler = new JwtSecurityTokenHandler();
            var token = handler.ReadJwtToken(accessToken);

            if (!token.Payload.ContainsKey("exp"))
            {
                return true;
            }

            var exp = token.Payload["exp"];
            if (exp is not long expTimestamp)
            {
                return true;
            }

            var expDateTimeOffset = DateTimeOffset.FromUnixTimeSeconds(expTimestamp).UtcDateTime;
            return DateTime.UtcNow >= expDateTimeOffset;
        }

        public static async Task<string> FetchToken(string currentToken, string clientId, string clientSecret, string url)
        {
            if (!string.IsNullOrWhiteSpace(currentToken))
            {
                if (IsTokenExpired(currentToken))
                {
                    return await GenerateTokenHttp(clientId, clientSecret, url);
                }
                else
                {
                    return currentToken;
                }
            }
            else
            {
                return await GenerateTokenHttp(clientId, clientSecret, url);
            }
        }
    }

    public class AccessToken
    {
        public string access_token { get; set; }
    }

}