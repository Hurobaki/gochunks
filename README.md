# Gochunks

This application allows you to create chunks from a list of files. You specify chunks size and it will creates folders with as many files as the size specified.


## Install

### Clone and install

```shell script
go get -u github.com/Hurobaki/gochunks
```

As any go application, it will download gochunks' sources to your `$HOME/go/src/github.com/Hurobaki/gochunks` folder.
You must add `$HOME/go/bin/` directory to your $PATH in order to be able to call any go executable to your path.

### Build it

```shell script
cd $HOME/go/src/github.com/Hurobaki/gochunks
go install -mod=vendor
```

## How to use

```shell script
gochunks <directory>
```

You can add `-zip` flag to get **.zip** output.

```shell script
gochunks -zip <directory>
```

And change the default chunk size

```shell script
gochunks -zip -size=10 <directory>
```

# Authors

*   **Théo Herveux** [MyGit](https://github.com/Hurobaki)