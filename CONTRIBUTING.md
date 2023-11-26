# Contribute to "prc"

Welcome to our open-source project! Your contributions are valued. Keep in mind that we rely exclusively on the 
official [procfs documentation][0] at kernel.org and the [proc(5)][1] man page. When contributing, ~~stick to these 
sources strictly~~. Avoid using information from other places.

# Step 1: Clone the source code

We encourage direct contributions through GitLab, our primary code hosting platform. However, if 
you find it more convenient, you can still submit pull requests from Github.

    % git clone https://gitlab.com/nzv/prc.git
    % cd prc
    % vi ...

> Follow "[Effective Go][2]" for contributions. Thanks!

# Step 2: Test your changes

Before submitting your code, conduct thorough testing using the provided Makefile. Run all tests in the project with:

    % make test

To test only a specific directory, modify the "DIR" environment variable. For instance, to test only the `tty` directory:

    % make DIR=tty test

# Step 4: Send changes for review

Follow Golang's commit message guidelines [here][3]. Our project aligns with these recommendations. Thanks!

    TODO

[0]: https://www.kernel.org/doc/Documentation/filesystems/proc.rst
[1]: https://www.kernel.org/doc/man-pages/online/pages/man5/proc.5.html
[2]: https://go.dev/doc/effective_go
[3]: https://go.dev/doc/contribute#commit_messages
