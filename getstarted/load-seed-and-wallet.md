--
title: Load Seed & Wallet
description: Load the seed into the controller software, and optionally include the wallet descriptor.
published: 2023-02-07
--

## Load Seed {#input}

All backups begin by loading the seed words you want to back up.

When backing up a multisig wallet, load any seed that is a part of the multisig setup.

- Choose "Backup Wallet" on the startup screen.
- Remove the SD card when prompted to do so.
- Select either keybord or camera input method.
    - Keyboard:
        - Choose number of words.
        - Input the seed words in the correct order.
    - Camera:
        - Scan a QR code containing a 12- or 24-word seed phrase (SeedQR, CompactSeedQR or Mnemonic).
- Confirm the seed.

<p class="alert alert-warning">Ensure the SD card is removed as prompted in the app before entering any seed.<br><br><img src="/static/img/remove-sd-card-warning.webp" alt="Remove SD card warning in the SeedHammer app"></p>

## Scan or Skip Output Descriptor/Public Key {#scan-output-descriptor-or-public-key}

Select either to scan, skip, or re-use an output descriptor or public key.

### Scan
Choose 'Scan' if you want the backup to include the public part of your wallet. This facilitates trustless loading into wallet or coordinator software, allowing for the viewing of balances and addresses.

We highly recommend this option, particularly for multisig wallets.

- Scan the output descriptor or public key (XPUB) of the wallet. Currently the software supports:
    - Output descriptors from [Sparrow Wallet](https://sparrowwallet.com/), [Nunchuk](https://nunchuk.io/), and [BlueWallet](https://bluewallet.io/).
    - Descriptors in text mode (raw or wrapped in JSON).
    - Raw public keys (XPUBs).
- Confirm the wallet.
    - Click the info-button to see the receive and change addresses of the wallet.

After confirming the wallet [start the engraving process](/get-started/engrave-plate).

### Skip
Choose 'Skip' if you wish to back up only the seed and nothing else.

You will immediately be directed to the [start of the engraving process](/get-started/engrave-plate).

<p class="alert alert-warning">If you skip scanning the output descriptor, you will be unable to restore a multisig wallet with only the minimum number of plates at hand (as described <a href="/article/how-to-recover-from-seedhammer-backup#multisig">here</a>), because there will be no descriptor side available.</p>

### Re-use
Choose 'Re-use' if you wish to re-use the last entered descriptor or public key from this session.

Note: Should the controller have been powered off, the option to re-use the descriptor will become unavailable.

### How to Scan a Descriptor
<details>
    <summary>Scan Output Descriptor from Sparrow Wallet</summary>
    <p>
        <ul>
            <li>Open the wallet you want to back up.</li>
            <li>Under "Setting"-tab click "Export" and choose export for either Jade Multisig, Keystone Multisig, Passport Multisig, or Specter Desktop.<br><img src="/static/img/sparrow-export-descriptor-with-title.webp" alt="Click the Export button in Sparrow"></li>
            <li>Scan the QR code.<br><img src="/static/img/sparrow-show-descriptor.webp" alt="Scan the QR into Sparrow"></li>
        </ul>
    <p>
</details>

<details>
    <summary>Scan Output Descriptor from Nunchuk</summary>
    <p>
        <ul>
            <li>Open the wallet you want to back up.</li>   
            <li>Under "View wallet config" -> More options menu (3 dots) -> "Export wallet configuration" click "QR code".</li>
            <li>Scan the QR code.</li>
        </ul>
    <p>
</details>

<details>
    <summary>Scan Output Descriptor from BlueWallet</summary>
    <p>
        <ul>
            <li>Open the wallet you want to back up.</li>
            <li>Under the options menu (3 dots) click "Export Coordinator Setup".</li>  
            <li>Scan the QR code.</li>
        </ul>
    <p>
</details>
