--
title: Install The Software
description: Installation and verification of the controller software.
published: 2023-02-07
--

## Download {#download}

Get the latest version of the <a href="https://github.com/seedhammer/seedhammer/releases/latest" target="_blank">SeedHammer software</a>.

- Download the disk image `seedhammer-vX.Y.Z.img`
- Download the signature file `seedhammer-vX.Y.Z.img.sig`
 
 ...where `vX.Y.Z` is the most recent version number.

In addition download our <a href="/static/seedhammer-signers.pub" download>public signing key</a>.
- This file should be downloaded only once and re-used for verification of future software versions.

All 3 files should be in the same directory:

<img src="/static/img/download-software-files.webp" alt="all files should be in the same directory" style="width:initial;margin-left:initial;border:1px solid #e1e1e1;">

## Verify Image {#verify}

Verify the software image signature.

- Verify the `seedhammer-vX.Y.Z.img` with the ssh-keygen tool:
<code>$ ssh-keygen -Y verify -f seedhammer-signers.pub -I gh@seedhammer.com -n seedhammer.img -s <strong>seedhammer-vX.Y.Z.img.sig</strong> < <strong>seedhammer-vX.Y.Z.img</strong></code>
- The command should succeed with a message similar to:
```Good "seedhammer.img" signature for gh@seedhammer.com with ED25519 key SHA256:xqo5VI7m8vh8UHgceOt27dep1EZ7SyCFjJOogcQBrvw```

<p class="alert alert-warning">
For extra security build the software from source code. Instructions are in our <a href="https://github.com/seedhammer/seedhammer#building-from-source" target="_blank">GitHub README</a>.
</p>

## Write Image {#write}

Write `seedhammer-vX.Y.Z.img` to an empty or unused microSD card.

On Windows and macOS, use a tool like [balenaEtcher](https://www.balena.io/etcher#download-etcher).

On linux, use a tool like dd:

```dd if=seedhammer-vx.y.z.img of=/dev/sdx conv=fdatasync bs=1M```

## Start The Controller {#start}

<p class="alert alert-warning">
Never connect the controller to a computer, use a regular power adapter.
</p>

Test that the microSD card works by starting the controller.

- Install the microSD card into a compatible controller (SeedSigner).
- Power on the controller by plugging a micro-USB cable into the **outermost port**.
- Remove the microSD card when the startup screen appears.

![Power device](/static/img/power-device.webp)

**Important:**

- The controller should be powered on via the outermost USB port as the SeedHammer machine **must** use the middle port.

After starting the software, proceed [to load a seed](/get-started/load-seed-and-wallet).
