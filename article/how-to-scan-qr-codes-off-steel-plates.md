--
author: SeedHammer
authorlink: https://twitter.com/SeedHammer
title: How to Scan QR Codes off Steel Plates
description: Learn the technique to easily read the QR codes from the SeedHammer steel plates.
published: 2023-04-07
image: scan-qr-of-steel.webp
--

The technique is demonstrated in this video and explained in depth below.

<div class="embed-responsive embed-responsive-16by9"><iframe class="embed-responsive-item" src="https://www.youtube.com/embed/SCTtZ-DvZ2Y" title="How to Scan a QR off Steel Plates"></iframe></div>

## Dark QR on Bright Background

Most simple QR code scanners expect the QR code to be dark on a bright background. Amongst others this goes for the SeedSigner, Blockstream Jade, Sparrow Wallet and the SeedHammer controller software itself.

![Light vs. dark QR on SeedHammer steel plates](/static/img/light-vs-dark-qr.webp)

As soon as the QR turns dark it is easy readable.

## How To Scan a Seed QR {#scan}

To obtain a dark QR on a bright background use a light source to hit the plate at an angle.

1. Position yourself so a wide/diffuse light source (a lamp, sun from a window etc) is behind you.
1. Hold the SeedHammer steel plate up like a small mirror, and pan/tilt the plate until the light source is reflected in the center of the QR code.
1. The QR code now turns dark and can be scanned.

<div class="grid">
    <div><img src="/static/img/scan-seedsigner.webp" alt="SeedSigner scanning a CompactSeedQR from a SeedHammer backup plate"></div>
    <div><img src="/static/img/scan-jade.webp" alt="Blockstream Jade scanning a CompactSeedQR from a SeedHammer backup plate"></div>
    <div><img src="/static/img/scan-seedhammer.webp" alt="SeedHammer app scanning a CompactSeedQR from a SeedHammer backup plate"></div>
</div>

## How To Scan a Descriptor QR

The theory is the same as described [above](#scan). However scanning of a descriptor QR can be more tricky because the QR details are smaller.

Furthermore a static mounted webcam is often used to scan a descriptor into the coordinator software, limiting pan and tilt maneuvering.

One way to do it is:

1. Activate the descriptor scan mode in the coordinator software (see [this guide](/get-started/perform-recovery#multisig)).
1. Make the screen as bright as possible (max brightness and white pixels on the screen) in order to make it function as light source.
1. Present the descriptor side of the plate (never the Seed side!) to the webcam.
1. Pan and tilt the QR until it turns dark and can be scanned.

![Sparrow Wallet scanning a descriptor QR from a SeedHammer multisig backup plate](/static/img/scan-sparrow.webp)

On Mac it is possible to use your iPhone as external camera which can make scanning easier.

## Increase The Readability of a QR

Increase the readability of QR on steel by painting the QR with a black marker and wipe off the residual.

This makes any QR about as readable as QR codes on paper or on a screen.

Read the full article on applying a [solid marker here](/article/enhance-qr-readability-with-solid-marker).

![QR code painted with a marker for higher readability](/static/img/qr-with-marker.webp)