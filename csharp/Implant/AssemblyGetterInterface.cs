using System;
namespace Implant
{
	public interface AssemblyGetterInterface
	{
        Task<string> GetBase64AssemblyAsync();
	}
}

