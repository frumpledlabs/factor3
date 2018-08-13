# Factor 3

#### Description
Opinionated environment config loading for Golang.  Intended to strictly follow Factor 3 of the [Twelve Factor App](https://12factor.net/) methodology.

#### Notes
This project is still in development (notice version 0).  While the interface to this package is intended to remain stable, that cannot be promised.  Count on this package still solving the problems illustrated in the examples, not necessarily the interface to use or initialize this package.

#### Currently planned improvements:
- Add logging (with option to inject your own logger).
- Add docs (godocs)
- Clean up the code base in general.  I found dealing with reflection in Golang to be quite a nightmare; currently, the codebase reflects this D:
