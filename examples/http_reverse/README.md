# Latest http example

This example show cases how you can run a http webserver with versionify. This example proposes a different structure where there is one latest version.

This "latest" version contains all the latest relevant code. Any lower versions extend on this latest version adding/overiding methods with what is still deprecated.
It is also called reverse, since we take the highest version (latest) and any lower version extends on this version.


This allows for a different way of working. You work primarily in the latest directory and when there is a breaking change you up the latest version and move the handler to its respective older version directory.