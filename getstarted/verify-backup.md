--
title: Verify Backup
description: Verify that you can trust the backup when you need to recover.
published: 2023-02-07
--

## Singlesig Backup {#singlesig}

Verifying a singlesig backup:

1. Verify that the master fingerprint is correct.
1. Verify that the seed words are correct (and in the same order).
1. Verify that the QR contains the same data by scanning it with a [SeedSigner](https://seedsigner.com/software/).
    - Boot up a SeedSigner.
    - Choose "Scan".
    - Confirm the fingerprint displayed on the screen.
    - Navigate to "Backup Seed" -> "View Seed Words".
    - Confirm the seed words.

## Multisig Backup {#multisig}

### Verify Seed Side

For each plate in the multisig backup:

1. Verify that the master fingerprint is correct.
1. Verify that the seed words are correct (and in same order).
1. Verify that the numbering (eg. 1/3, 2/3, etc.) is as expected.
    - No two plates should have the same numbering unless you have made a copy of that particular plate.
1. Verify the content of the CompactSeedQR by scanning it with a stateless signer.
    - Ask the signer to show the backup seed.
    - Verify that the seed words shown on the signer are correct (and in same order).

### Verify Descriptor Side

1. Verify that the QR code contains the text written next to it.
    - Scan it with an offline QR scanner (eg. QR scanner app on a mobile phone with flight mode enabled).
        - Be aware that some QR scanners save a list of recently scanned QR data. Remember to delete this list and delete the app before going out of flight mode again.
1. Perform a test recovery with any combination of the quorum.
    - Create a new wallet in Sparrow Wallet.
    - Under "Script Policy" click the camera icon. ![Click Script Policy in Sparrow](/static/img/sparrow-import-descriptor.webp)
    - Scan the QR code(s) on the side **that does not** contain the seed words. ![Scan the QR into Sparrow](/static/img/sparrow-scan-descriptor.webp)
    - In a 2-of-3 multisig recover with any of these combinations of plates:
        - Share 1 and 2.
        - Share 1 and 3.
        - Share 2 and 3.
    - In a 3-of-5 multisig recover with any of these combinations of plates:
        - Share 1, 2, and 3
        - Share 1, 2, and 4
        - Share 1, 2, and 5
        - Share 1, 3, and 4
        - Share 1, 3, and 5
        - Share 1, 4, and 5
        - Share 2, 3, and 4
        - Share 2, 3, and 5
        - Share 2, 4, and 5
        - Share 3, 4, and 5

**Important:**

<p class="alert alert-warning">Never bring the seed QR codes near the camera controlled by the Sparrow Wallet.</p>

<p class="alert alert-warning">Never input the seed words into Sparrow Wallet.</p>

In the case of singlesig we recommend to [perform a test recovery](/get-started/perform-recovery).