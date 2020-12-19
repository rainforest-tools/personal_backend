title: 20200812 Gyroid Nacre Modeling AI 3DP
tags: Research, Internal Meeting
description: 
---

# Gyroid Nacre Modeling AI 3DP

[![hackmd-github-sync-badge](https://hackmd.io/gxT36hquRu2s9TJWUTwHZg/badge)](https://hackmd.io/gxT36hquRu2s9TJWUTwHZg)


---

## [BOTORCH](http://arxiv.org/abs/1910.06403)
![](https://i.imgur.com/eyv8lwC.png)

----

```python
import torch
from botorch.models import SingleTaskGP
from botorch.fit import fit_gpytorch_model
from botorch.utils import standardize
from gpytorch.mlls import ExactMarginalLogLikelihood

train_X = torch.rand(10, 2)
Y = 1 - torch.norm(train_X - 0.5, dim=-1, keepdim=True)
Y = Y + 0.1 * torch.randn_like(Y)  # add some noise
train_Y = standardize(Y)

gp = SingleTaskGP(train_X, train_Y)
mll = ExactMarginalLogLikelihood(gp.likelihood, gp)
fit_gpytorch_model(mll)
```

---

## Architecture
### Data
```mermaid
graph TB
    subgraph Data
        subgraph Storage
            StorageNode(" ")
            GoogleDrive("Google Drive") --&gt; StorageNode(" ")
            Local --&gt; StorageNode(" ")
        end
        subgraph Version Control
            Git
        end
        StorageNode(" ") --&gt; DVC
        Git --&gt; DVC
    end
    subgraph Algorithm
        subgraph Kubeflow
            JupyterServer("Jupyter Server")
            Pipeline
        end
        DVC --"Development"--&gt; JupyterServer("Jupyter Server")
        DVC --"Development"--&gt; VSCodeServer("VS Code Server")
        JupyterServer("Jupyter Server") --"Training"--&gt; Pipeline
        VSCodeServer("VS Code Server") --"Training"--&gt; Pipeline
    end
```

----

#### [DVC](https://dvc.org/)
* [Supported Storage Types](https://dvc.org/doc/command-reference/remote/add#supported-storage-types)
    * Google Drive
    * Local Remote
    * SSH
    * HTTP
    * WebDAV
    * Google Cloud Storage
    * Amazon S3
    * ...

----

### Development
```mermaid
graph TB
    subgraph Data
        DVC
    end
    subgraph DevOps
        GithubRepo --"CI/CD"--&gt; DockerHub("Docker Hub")
    end
    subgraph Algorithm
        subgraph Kubeflow
            JupyterServer("Jupyter Server")
            Pipeline
        end
        DVC --"Development"--&gt; VSCodeServer("VS Code Server")
        DVC --"Development"--&gt; JupyterServer("Jupyter Server")
        JupyterServer("Jupyter Server") --"Training"--&gt; Pipeline
        VSCodeServer("VS Code Server") --"Training"--&gt; Pipeline
        Pipeline --"Hyperparameters Tuning"--&gt; Katib
        
        DockerHub("Docker Hub") --&gt; JupyterServer("Jupyter Server")
        DockerHub("Docker Hub") --&gt; Pipeline
        DockerHub("Docker Hub") --&gt; Katib
    end
    Pipeline --"Sending Jobs"--&gt; ComputingResources("Computing Resources")
```

----

### DevOps
```mermaid
graph TB
    subgraph Data
        DVC
    end
    subgraph DevOps
        GithubRepo --"CI/CD"--&gt; DockerHub("Docker Hub")
    end
    subgraph Algorithm
        subgraph Kubeflow
            JupyterServer("Jupyter Server")
            Pipeline
        end
        DVC --&gt; VSCodeServer("VS Code Server")
        DVC --&gt; JupyterServer("Jupyter Server")
        JupyterServer("Jupyter Server") --&gt; Pipeline
        VSCodeServer("VS Code Server") --&gt; Pipeline
        Pipeline --&gt; Katib
        
        DockerHub("Docker Hub") --"Supplying Images"--&gt; JupyterServer("Jupyter Server")
        DockerHub("Docker Hub") --"Supplying Images"--&gt; Pipeline
        DockerHub("Docker Hub") --"Supplying Images"--&gt; Katib
    end
```
