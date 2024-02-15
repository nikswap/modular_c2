using System;
namespace Runner
{
	public class RunnerClass
	{
		public static string DoIt()
		{
			Console.WriteLine("USER: " + Environment.UserName);
            return "Hello from DoIt";
		}
	}
}

