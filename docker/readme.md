barakb/create-branch
====================

A docker file that create image that serve webapp that can create or delete git branches from multiple repositories.

### Downlading:

`docker pull barakb/create-branch-docker`


### Getting Github client id and secret key

Go to [https://github.com/settings/applications/new] and create new application. 
In the callback, put the application url with /githuboa_cb as the callback URL (e.g. https://IP:PORT/githuboa_cb).

Once finished, it will show client id and secret id. Save those for next step.

### Configuring:

create a conf dir named 'conf'  with the following files:

`conf/env.sh` with the following conetnt

```bash
  export ClientID="github client id"
  export ClientSecret="github client secrete"
```

`conf/repos.txt`

 A file with one line for each git repository in the format owner/reponame for example:

`barakb/foobar`



### Running

`./run.sh`
