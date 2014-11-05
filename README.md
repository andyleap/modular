modular
=======

Modular GoLang System

modular allows you to create "on demand" style module loading easily.  
This is useful both for on demand module loading, and for decoupling modules from a main program.


It's a very simple wrapper around the net/rpc package.
To create a module, just make a new command, import the module package, and do

    module := module.New()

Using standard rpc functions, register your type, then `module.Serve()` to start the module.
`module.Serve()` returns when the module is unloaded

To use a module, import the loader package, and do

    module := loader.load("module", "arg1", "arg2")

At this point, module is a RPC client, and you can call functions remotely.
When done with a module, call `Close()` on it, and the module will close.
