# 🚀 Contract Testing with Pact Go

## Introduction

> This project offers an basic introduction how to create a simple contract testing with `Pact Go`.

> Learn everything `Go` language in [here](https://go.dev/doc/tutorial/getting-started#code)

## Prerequisites

### ***Step 1:***

- Install go https://go.dev/doc/install

### ***Step 2:***

- Install [Pact](https://github.com/pact-foundation/pact-go#installation)

#### 1️⃣ Download the Latest Version

Pact:  [Pact Ruby Standalone Releases](https://github.com/pact-foundation/pact-ruby-standalone/releases)
> 👉 For macOS/Linux → Download pact-<version>-osx.tar.gz or pact-<version>-linux-x86_64.tar.gz

> 👉 For Windows → Download pact-<version>-win.zip

#### 2️⃣ Extract and Install

- Extract the file (replace <version> with the actual version number):

```bash
tar -xzf pact-<version>-osx.tar.gz -C /opt
```

- Add Pact to your PATH

```bash
echo 'export PATH=$PATH:/opt/pact/bin' >> ~/.bashrc
source ~/.bashrc  # (or use ~/.zshrc for macOS zsh)
```

### ***Step 3:***

- Install `Pact FFI`: https://docs.pact.io/implementation_guides/rust/pact_ffi

> The `libpact_ffi` library refers to the native shared library that allows Go (or other languages) to interface with
> the
> `Pact FFI` (Foreign Function Interface) bindings. It's a key component for making Pact-based contract testing work
> when
> using native bindings.

> Here’s how to find and correctly reference the `libpact_ffi` library:

#### 1️⃣ What is `libpact_ffi`?

The `libpact_ffi` is the compiled shared library created from the `Pact FFI Rust` bindings. It exposes the Pact
functionality (such as contract verification) through C-compatible functions, which can be used by Go and other
languages.

#### 2️⃣ Where to get `libpact_ffi`?

You can get the `libpact_ffi` library by building it from the `Pact FFI Rust` source code:

#### 3️⃣ Setup `libpact_ffi` library:

3.1 Clone the Pact Reference Repository:

```bash
git clone https://github.com/pact-foundation/pact-reference.git
cd pact-reference/rust/pact_ffi
```

3.2 Build the FFI Library:

- To build the library from source, you need to use cargo, the Rust package manager. First, make sure you have Rust
  installed. You can check by running:

```bash
rustc --version
```

- If it's not installed, you can install Rust using:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

- After Rust is installed, you can build the library using:

```bash
cargo build --release
```

This will compile the `libpact_ffi library` and place it in the `target/release` directory of the pact_ffi
project.

3.3 Move the Library to /usr/local/lib:

- Now that you've built the `libpact_ffi.dylib` file, move it to `/usr/local/lib` to make it available for use
  by your system. You will need to use sudo to have the necessary permissions:

```bash
sudo mv /path/to/pact/ffi/libs/libpact_ffi.dylib /usr/local/lib/
```

3.4. Update the Library Path (Optional):
If the library is in `/usr/local/lib`, you may also need to update the `DYLD_LIBRARY_PATH` to ensure that your
system can find it:

```bash
export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:/usr/local/lib
```

3.5 Verify the Configuration

- After configuring `DYLD_LIBRARY_PATH`, you can verify that it's set correctly by running:

```bash
echo $DYLD_LIBRARY_PATH
```

This should print the current value of the `DYLD_LIBRARY_PATH` variable, showing the directories you've configured.

## Project structure

## Steps to perform Contract Testing

### 1 Generate/update a contract:

1. cd /client
2. go test

### 2 Run provider test to validate it against the contract:

1. cd /server
2. go test

## Reference

https://docs.pact.io/ <br>
https://github.com/pact-foundation/pact-go <br>
https://docs.pact.io/implementation_guides/go <br>
https://www.linkedin.com/pulse/learnings-observations-from-implementing-contract-testing-herring/ <br>
https://www.youtube.com/playlist?list=PLwy9Bnco-IpfZ72VQ7hce8GicVZs7nm0i

