--
title: How It Started
description: What problem does SeedHammer solve for Bitcoiners and how an innovative new way of securing multisig backups was invented.
--

The story behind SeedHammer starts in 2021, when we were searching for an easy, secure and scalable way to create backups of multisig bitcoin wallets onto a durable material like steel. 

While there were lots of manual metal backup options available, they were either missing a way to store the wallet descriptor or XPUBs, had security or reliability flaws or were tedious and time consuming to use when backing up complex multisignature wallets.

## XPUBs Must Be Part of the Backup

For multisig setups especially, just backing up the seed phrases of each share is not enough. Without also having the XPUBs for all the seeds in a multisig wallet, you can’t later recover the wallet or move the funds.

A so-called wallet “output descriptor” contains exactly these XPUBs (and a bit more information) in order to recreate a multisig wallet fast and easy. Most wallets recommend separately backing up the wallet descriptor on paper or using online storage, but that’s either fragile or it risks your wallet’s privacy, or both. 

You want to store those XPUBs as securely and safely as your seed words, but none of the existing metal seed backup solutions catered for this. We ended up with the good old embossing stamps and a bunch of steel plates from the local hardware shop.

## A Painful Experience

Individually hammering the descriptor (or all the XPUBs) by hand quickly became unmanageable, tedious and physically unpleasant. Furthermore, the result was not convincing in its visual appearance, and it was hard to validate the accuracy of the long strings of text. A single error while hammering a plate means throwing the plate away and starting again.

Loading the XPUB backups back into the wallet coordinator software for testing was also extremely tedious and error prone. The sleep-well-at-night factor was simply not high enough.

Due to these impracticalities we looked into the possibility of developing an automated solution that not only would do the process faster and easier but also introduce the possibility for easy recovery with QR codes.

## An Automated Solution

We believed the ideal solution was to have a provably air gapped offline “hammer” that can automatically engrave all the required information, both the seed and the descriptor onto a set of metal plates. 

In the search for an automated backup solution, it became clear that there was no existing solution available that was specific to a Bitcoiner’s needs. After months of research and trial and error, we finally found an off-the-shelf industrial engraving machine that was:

- durable enough for the job of hammering multisig setups.
- provably air gapped and stateless.

We then added some improvements to the physical machine and developed a simple, intuitive and easy-to-use controller software. Together we called it the SeedHammer.

## The Controller {#controller}

We started to develop the controller on the same development board as Specter DIY. However, mainly because of delivery problems, we chose to move away from this approach.

In frustration we chose to build the software as an app that could run on Android, Mac OS, Linux etc. letting the user decide which unit to run it from - and if the user wanted a controller with the machine we would offer a cheap burner phone.

But we never felt quite satisfied with this approach as it would be too tempting to run it on a device that would not be truly air gapped. Furthermore we could not satisfactorily prove that the phone would not leak the private keys.

Yet again we switched platform and chose to use the same hardware setup as SeedSigner (Raspberry Pi Zero). A few things was in favor of this choice:

- Around half the price of a Specter DIY (incl. case) at that time.
- The Specter board is a development board (no guarantee for its continuation).
- The obvious advantage of Specter DIY, secure boot, can be circumvented.
- An statement from The Raspberry Pi Foundation that Pi Zero soon will be publicly available again in satisfying quantities.

Without any evidence we also had a feeling that many of our initial users would own or be familiarized with a SeedSigner already. And the choice did not rule out a later version of the software that could run on any other open source signing device. 

## Self-contained Setup

Without violating any industry standards we wanted to figure out a way so the minimum amount of physical backup steel plates would be enough to recover the full multisig wallet - without a separate descriptor stored anywhere.

We did explore the idea of adding the full descriptor on the backside of each plate, but it would take far too long for the machine to hammer. Furthermore having the full descriptor on every plate is also a privacy risk as anyone with access to just a single plate would be able to easily see the full transaction history.

Then it dawned on us – why not split up the descriptor into its component parts and distribute those parts rotationally onto the back of each of the seed plates. This way, any quorum of plates would inherently contain both the minimum number of seeds required to sign transactions as well as all the information needed to recreate the full wallet.

To give the user further comfort, we add onto the seed plates some additional non critical information like the fingerprint, title as well as the number of seeds (M-of-N).

## Recover With Minimum of Shares

The new way to structure the data on each plates front- and backside makes the minimum amount of plates enough to recover the wallet - even without any separately stored full descriptor.

Another security benefit is that the wallet balance and transaction history is also only easily available to people who have physical access to the M-of-N signing quorum – so only those able to move coins can easily see how many coins are in the wallet.

In all other multisig backups we looked at on the market, losing the wallet descriptor and just a single plate (and therefor the derived XPUB) would mean an irretrievable loss of access to the funds in the multisig wallet.

This problem is now solved with the SeedHammer.