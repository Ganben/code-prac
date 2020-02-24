# - Julia version: 1.1.1
#
# args error
module MyArgs

include("args.jl")

end

name, _ = MyArgs.get_parameter("name", true)
if name == "":
    name = "handsome"
end
println("hey julia! $(name)!")