# jenkins-from-scratch

Aims to remove any cyclical or external dependencies from low-level platform/infra deployments. Since these
are the lowest unit in a platform orchestration tree. It's imperative that it has no cyclical, external or network dependencies.

Manages Jenkins completely from code, and allows it to be maintained, upgraded & recovered from
the command-line, using no external system dependencies other than EC2.

In an [outage](https://substack.com/history/post/168964142), security compromise or other such SNAFU - critical platforms need to be independent of any other platform/system which may be unavailable.

In other words - your lowest level bootstrapper cannot rely on anything other than itself. It needs to be completely self-contained.

If you think that sort of thing never happens...

[It](https://engineering.fb.com/2021/10/05/networking-traffic/outage-details/) [sure](https://blog.bytebytego.com/p/how-the-google-cloud-outage-crashed) [does](https://www.gremlin.com/blog/the-2017-amazon-s-3-outage). And when it does it has the potential to take down a business or at least cost $$$ in downtime & reputation.

jenkins-from-scratch delivers a Jenkins server managed purely from the CLI with ALL configuration stored in the codebase. It doesn't rely on backups or any higher-level platform/service.

You can use this Jenkins to spawn more for niche services, or bootstrap a datacenter, or Terraform a cloud account, and you can be sure that it can always work as long as you have a copy of the codebase.

Additionally - Jenkins (in my experience) is a great tool but usually poorly implemented. It's old-school Java, but that doesn't mean managing it has to be a pain. This project handles backups, recovery, rollbacks, rebuilds, plugin versioning, state management issues and all the other headaches that drive many away from using Jenkins. If you need a robust, tried-true, self-hosted, highly flexible & capable job runner, then it's hard to look past Jenkins IMHO.

# Using it
Install Poetry and resolve project dependencies
```bash
brew install poetry
poetry install
```

## Generate a new keypair locally
This is uploaded to EC2
```bash
poetry run task create-ssh-key
```

## Get the Jenkins password after initial startup
Ansible will display the initial admin password. This only required the first time the datastore is built.

## Deploy it
Deploys all resources and connects to the instance using the Ansible dymanic inventory plugin
```bash
poetry run task deploy
```

## Testing
```bash
poetry run task lint
poetry run task test
```

# Features

## Plugins from Code
Define plugins in the `plugins.txt` file using the `<plugin_name>:<plugin_version>` format.
To always pull the latest version of a plugin, ommit `<plugin_version>`

These are baked into the jenkins image build with each deployment.

## System Configuration As Code CASC
Uses the [CASC plugin](https://plugins.jenkins.io/configuration-as-code/) to maintain system config as YAML in the codebase

## Job Configuration As Code
TODO:
XML based job configs are managed in the codebase.

## SSL Proxy
TODO: Add ALB to front Jenkins UI

## Extensible Agents
TODO:
Add as many agents as you like, or build on the main node

## Job History Persists Rebuilds
TODO:
External volume attachment

## Snapshot based backups
TODO:
This is available to back-up the entire directory

## Docker based
Uses the `jenkins/jenkins:lts-jdk17` public image from https://hub.docker.com/r/jenkins jenkins as a base image

Builds and deploys happen on the same host.

## Secrets Manager (SSM) Integration
TODO:
Using the secrets manager plug-in, automatically presents remote secrets as credentials