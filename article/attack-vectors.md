--
author: SeedHammer
authorlink: https://twitter.com/SeedHammer
title: Attack Vectors
description: Known security vulnerabilities to using SeedHammer.
published: 2023-02-25
image: evil-attack.webp
--

Your seed phrases are your bitcoin, so it is wise to be hypervigilant of the risks of using any electronic device that will come into contact with your seeds. 

The risk we are defending against is exfiltration of seed phrases that the SeedHammer has come into contact with – this is the objective any attacker will want to achieve.

Here are the attack vectors that we are working on diligently defending you from.

## Sound {#sound}

With sufficient knowledge of how a machine like SeedHammer works, being in close proximity, recording sound at the time actual seeds are being hammered, there is a theoretical possibility that someone could figure out the seed words being hammered into your plates.

<del>So when you do setup your SeedHammer, for now, think about doing so in a location where your machine cannot be easily overheard. Consider unplugging any smart speakers, powering down mobile phones and any device capable of recording sound in the area. You may also wish to play loud recorded music while the hammer is running.</del>

Since v0.11.1 of the SeedHammer controller software exfiltration of the private data is addressed by the introduction of [constant time engraving](https://github.com/seedhammer/seedhammer/releases/tag/v0.11.1):

<div class="embed-responsive embed-responsive-16by9"><iframe class="embed-responsive-item" src="https://www.youtube.com/embed/b3gb3Cgpwnc" title="Elimination of Sound Exfiltration: Seed Phrase"></iframe></div>

This also applies to the SeedQR codes:

<div class="embed-responsive embed-responsive-16by9"><iframe class="embed-responsive-item" src="https://www.youtube.com/embed/jIbB9UYTz4w" title="Elimination of Sound Exfiltration: Seed Phrase"></iframe></div>

## Radio Waves

Systems that use Bluetooth or WiFi transmit data that can be potentially decrypted and eavesdropped. That’s why SeedHammer has been designed around only using components that have no wireless or network transmission capabilities at all.

The entire system, from the hammering machine to the Raspberry Pi Zero controller, is completely air-gapped from the internet.

## Electromagnetic Noise

The stepper motors that control the engraving process on the SeedHammer produce very faint electromagnetic waves. A very sophisticated eavesdropper in close proximity to the machine - when it is operating - could potentially figure out a way of using those electromagnetic waves to guess what the machine is hammering into the plates.

It's obviously an unlikely scenario. A simple precaution – don’t tell anyone when or where your SeedHammer machine is going to be used.

## Data Exfiltration {#data-exfil}

The SeedHammer app can potentially exfiltrate your seeds by storing them on the SD card in the Raspberry Pi Zero. This can be due to either:

1. The image loaded onto the SD card has been tampered with by a middleman.
2. The original open source code contains an undiscovered backdoor.

To ensures that you truly have the image we intended to distribute always [verify](https://seedhammer.com/get-started/install-the-software#verify) the image you download from our [Github](https://github.com/seedhammer/seedhammer). 

To eliminate any risk what so ever for any data exfiltration follow the in-app instructions and simply remove the SD card before entering any seeds.

![Remove SD card warning in the SeedHammer app](/static/img/remove-sd-card-warning.webp) 

## Onlookers

Don’t invite friends over to show them your SeedHammer hammering your actual seeds. Also, keep your device away from windows without blinds, security cameras or phones, laptops, TVs and tablets when hammering your seeds.

## Supply Chain Attack

The more obvious it is that a device is used to secure bitcoin, the more likely its supply chain will be of interest to criminals.

The Seedhammer machine and the Raspberry Pi components that make up the SeedSigner controller are all non-bitcoin specific devices.

We’ve also designed the architecture of the SeedHammer system so that you don’t have to trust the supply chain it comes from.

## Auto Generation of Seed Words {#generate-seed}

Bitcoin specific hardware products often try to help speed up the user experience by automatically generating a ready-to-use random seed phrase.

Because of the below mentioned risks associated with auto-generation of seeds, we do not offer auto-generation in the SeedHammer controller software:

- Backdoors

    Any software or hardware wallet that offers to generate bitcoin private keys should be considered dangerous. It is impossible to rule out the risk of a backdoor in the software. Even if the company that makes the wallet is reputable, the device running the software could be compromised by third parties.

- Pre-loaded Seeds

    Even air gapped devices running software that can produce private keys should not be considered 100% safe.

    By air gapping on specialised hardware, it may be impossible for a third party to exfiltrate the keys, but the firmware the hardware came with may have been preloaded with keys already known to a malicious third party.

    Some signers and wallets offer the user the option to introduce their own entropy via dice rolls or other techniques in order to improve your trust in the seed generation process. However, it is hard to prove conclusively that the software is actually using the entropy the user has introduced.

- Weak Entropy

    Even if the software is open source and the developers behind a project are well intentioned, truly strong entropy is notoriously hard to produce.

We encourage our users to think carefully about how their seed phrases they are backing up were generated and the risks involved in not generating your own seed phrases from scratch.

*In order to generate your seeds as secure as possible consider using coin flips and/or dice rolls to manually generate strong entropy. There are many video tutorials available online to teach you the simple steps needed to create your own strongly random seeds entirely offline and without trusting a third party.*

## Comments

Do you have any questions, comments or maybe ideas on how to improve the security of SeedHammer do not hesitate to contact us at [team@seedhammer.com](mailto:team@seedhammer.com).