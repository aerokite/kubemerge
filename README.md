## kubemerge

If you have multiple *kubeconfig*, you can merge them using **kubemerge** command.

### Install

```shell
$ wget https://github.com/aerokite/kubemerge/raw/master/bin/kubemerge
$ chmod +x kubemerge
$ mv kubemerge /usr/local/bin/
```

### Usage

If you want to merge your new *kubeconfig* with existing `~/.kube/config`, run following command

```shell
$ kubemerge ~/Downloads/kubeconfig-new
```

> Here, this `~/Downloads/kubeconfig-new` is your new *kubeconfig* file.

```shell
Success: merge completed
--> To restore your previous kubeconfig

	cp ~/.kube/config-20190404T004627 ~/.kube/config
```
