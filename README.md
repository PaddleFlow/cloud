# Cloud
Cloud repository provides APIs to easily debug PaddlePaddle in the local environment, and run distributed training in the cloud.

## PaddlePaddle on Baidu Cloud Tutorial

### Installation Requirements: 

* Python
* Authenticated Baidu Cloud account
* BaiduCloud IAM
* Authenticated Container Registry account

### Install latest release

```
pip install -U paddleflow-cloud
``` 

### Install from source

```
git clone https://github.com/paddleflow/cloud.git
cd cloud
pip install src/python/.

```

## High Level Overview

After debugging PaddlePaddle code on your local machine, you can easily run same code on BaiduCloud by using paddleflow cloud.

```
import paddleflow_cloud as pfc
pfc.run(entry_point='paddle_train_example.py')

```

## Setup Instruction

Follow the [instruction](/docs/setup_instruction.md) to setup the environment.

## User Guide

Follow the [User Guide](/docs/user_guide.md) for complete examples.

## Contributing

Follow the [Contributing](/docs/contributing.md) for more information of contribute to this repository.

## License

[Apache License 2.0](/LICENSE)

