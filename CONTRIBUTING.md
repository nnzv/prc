# Contribute

We talk about the project, fix things, share news, and ask questions on [GitLab]. There's also a mirror on 
[GitHub], but that's just for looking at, not for talking. Note that we depend solely on the [proc(5) man 
page]. You can also use other sources mentioned in the "SEE ALSO" section of that man page. When contributing, 
strictly follow these sources. Do not use information from elsewhere.

# Style guide

Refer to the [Go Code Style Guide] for guidance on the recommended layout of Go code and some 
style guide notes. Additionally, the [Google Go Style Guide] is a useful resource for 
understanding and crafting high-quality source code.

# Reporting bugs

To report bugs, use the "Bug" template in the [GitLab] repository. When using this template for an issue 
related to a feature or change, CC the responsible developers if known. This increases the chance of 
them reviewing and commenting. If you're uncertain about the developers, consult the git logs for the 
necessary information.

# Check and Test

When modifying Go code in the project, use the tests and checks in the build system, facilitated by the simple tool `run.go`.

1. To make things easier, you can use this alias:

       alias run="go run run.go"

2. Before testing, check your changes using the "run.go" script with the "check" target. This 
   checks for any issues in your code and ensures it follows the rules.

       vi kernel/uptime.go  # Make your changes
       run check

3. If everything is okay, run tests to be sure. This makes sure 
   your changes don't break anything. 

       run test

[GitLab]: https://gitlab.com/nzv/prc
[GitHub]: https://gitlab.com/nnzv/prc
[proc(5) man page]: https://www.kernel.org/doc/man-pages/online/pages/man5/proc.5.html
[Go Code Style Guide]: https://go.dev/doc/effective_go.html
[Google Go Style Guide]: https://google.github.io/styleguide/go/guide
