# Profiles in `kubectl debug`

The `debug` subcommand in `kubectl` supports a flag called `--profile`. As per [KEP-1441](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cli/1441-kubectl-debug/README.md#debugging-profiles), the plan is to have several different profiles available that configure ephemeral containers spawned by `kubectl debug`.

As this talk required `securityContext.privileged=true` for the ephemeral container hosting `dlv`, I initially wrote a quick patch to the subcommand to allow for a `sysadmin` profile. Since this isn't cleaned up at all, I have not submitted it upstream yet. As the talk showcases the power of a privileged ephemeral container, I plan to help with upstream implementing profiles.

## kubectl-ephemeral

To achieve the functionality I needed for the talk, I wrote a - less ergonomic but more powerful - `kubectl` plugin called [`kubectl-ephemeral`](https://github.com/embik/kubectl-ephemeral). It's pretty simple and does not have a lot of functionality besides ephemeral container configuration, but it helped as stopgap solution.
