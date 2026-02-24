# Exercise 0: Environment and Libraries

---

## Objective
Set up the environment for Python development and verify installation with required libraries.

---

## Step 1: Virtual Environment Setup

### Create Virtual Environment
```bash
python3 -m venv venv
or
$ python -m venv prompt_env
```

### Activate Virtual Environment
```bash
source venv/bin/activate
or
$ source prompt_env/bin/activate
(prompt_env) $


## Step 2: Python Version Verification

```bash
(prompt_env) $ python --version
Python 3.12.3


## Step 3: Install Required Libraries

```bash
(prompt_env) $ pip install jupyter numpy pandas
Collecting jupyter
  Downloading jupyter-1.0.0-py2.py3-none-any.whl
Collecting numpy
  Downloading numpy-1.26.4-cp312-cp312-linux_x86_64.whl
Collecting pandas
  Downloading pandas-2.2.0-cp312-cp312-linux_x86_64.whl
Installing collected packages: jupyter, numpy, pandas
Successfully installed jupyter-1.0.0 numpy-1.26.4 pandas-2.2.0


## Step 4: Verify Imports

### Test Script (`test_imports.py`)
```python

# test_imports.py
import jupyter
import numpy
import pandas

print("All imports successful!")
print(f"NumPy version: {numpy.__version__}")
print(f"Pandas version: {pandas.__version__}")
```

### Execution
```bash
(prompt_env) $ python test_imports.py
All imports successful!
NumPy version: 1.26.4
Pandas version: 2.2.0
```


## Step 5: Hello World Program

### Code
```python

import jupyter
import numpy
import pandas

print("Hello, World!")
```

### Execution
```bash
(prompt_env) $ python hello.py
Hello, World!
```

---

