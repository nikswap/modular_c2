open System.Reflection


// Dynamic loading is based on: https://www.jamesdrandall.com/posts/compiling_and_executing_fsharp_dynamically_at_runtime/

// For more information see https://aka.ms/fsharp-console-apps
printfn "Hello from F#"

let getMemberInfo (name:string) (assembly:Assembly) =
  let fqTypeName, memberName =
    let splitIndex = name.LastIndexOf(".")
    name.[0..splitIndex - 1], name.[splitIndex + 1..]
  let candidates = assembly.GetTypes() |> Seq.where (fun t -> t.FullName = fqTypeName) |> Seq.toList    
  match candidates with
  | [t] ->
    match t.GetMethod(memberName, BindingFlags.Static ||| BindingFlags.Public) with
    | null -> Error "Member not found"
    | memberInfo -> Ok memberInfo
  | [] -> Error "Parent type not found"
  | _ -> Error "Multiple candidate parent types found"

//TODO Look at error handling in F#
let loadPlugin (name:string) (path:string) =
    let assembly = Assembly.LoadFile(path)
    Ok assembly


