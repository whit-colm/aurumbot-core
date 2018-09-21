# Aurum
*Does everything you need it to, and nothing you don't*.

Aurum is a framework for creating and hosting your own custom discord bot with expandable third-party plugins. `aurumbot/core` is the end product itself, the rest being extraneous libraries.

## Features

- Plug'n'play design means next to no programming skill required *if you're willing to read this*
- Compiled in go to run at lightning speed with few resources
- Expandable third party plugins
- Easy to host yourself for free. No patreon for "full features" or hosting fees.
- Pseudo-daemon via tmux.

## Installing and Running Aurum

[Please see the wiki](https://github.com/aurumbot/core/wiki)

## Installing Plugins

1. Download the plugin from wherever, place it wherever.
2. Run `goaurum mkplugin $PLUGIN_PATH` to compile the plugins into `.so` files. ($PLUGIN_PATH is optional, left blank it will be set to `./`.)
3. Place the new .so plugins into the `$BOTNAME/plugins` directory
4. Reload the plugins
	- Run `!reloadplugins` (requires discord server administrator perms)
	- Restart the bot

## Contributing to the Project

Contributions would be greatly appreciated, however to keep everything orderly and organized please make sure of the following:
- **make all requests on the development branch.** I will not accept anything to the master branch.
- Test your changes, on multiple systems if you can manage it. (For reference, changes are usually tested on blank DigitalOcean images of the 5 provided distros (default version) and a Mac.)

All contributions are made under the Apache 2.0 License. Thank you for your understanding and contribution to Aurum!

## Contributors and Special Thanks:

- The [Discord Gophers server](https://discord.gg/PcAaUG8), without the help and patience of users there this project would never have made it this far.

- [Rob Muhlestein](https://github.com/robmuh)
- [Gabe Miller](https://github.com/gabeart10)
- [Mark Muchane](https://github.com/muchanem)

# EOF