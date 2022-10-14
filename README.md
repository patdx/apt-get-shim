# apt-get-shim

This is a simple shim for the apt-get binary than can be used in Fedora with DNF.

It is very basic. I was trying to install Playwright on Fedora but ran into [this issue](https://github.com/microsoft/playwright/blob/1cedd805ed08600e87bf8a1a37f8d11ac38c9630/packages/playwright-core/src/server/registry/dependencies.ts#L90):

> Cannot install dependencies for this linux distribution!

The Playwright linux install scripts seem to only use `apt-get`. After I added this shim I could complete the install.
