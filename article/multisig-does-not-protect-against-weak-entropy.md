---
author: SeedHammer
authorlink: https://twitter.com/SeedHammer
title: Multisig Does Not Protect Against Weak Entropy
description: Read about the challenges of true randomness and how multisig doesn't guard against weak entropy sources.
published: 2023-11-28
image: randomness.webp
---

This article is inspired by the common misconception expressed on social media that multisignature (multisig) wallets protect against weak entropy.

In a [recent case](https://twitter.com/mononautical/status/1728946782611366105), someone allegedly sent 140 BTC to a fresh cold wallet, which was instantly swept by a third party. The attacker managed to steal almost 56 BTC while paying 83.7 in transaction fees. Since the wallet was emptied immediately, it must have been monitored - one very likely explanation could be weak entropy.

Many are of the belief that if those coins had been sent to cold storage in a multisig wallet, then this would never have occurred. But that is wrong and a misconception. 

<article>Multisig setups are not mitigating the danger of keys created with weak entropy.</article>

A crucial aspect often overlooked is the quality of entropy used in generating the private keys in multisig setups. The reality is simple: a multisig wallet consisting of shares all made with equally low entropy does not protect against a wallet hack.

## The Challenge of True Randomness
Generating truly random numbers is a significant challenge in cryptography. Computers struggle to produce real randomness, which is vital for secure cryptographic keys. This randomness, or entropy, is the cornerstone of a secure Bitcoin private key - no matter if that key is intended to be used for singlesig wallet or as a cosigner in a multisignature wallet.

A Bitcoin private key is essentially a random number. If the process of generating this number is predictable, the key becomes easier to guess or brute-force. Such vulnerabilities can be exploited by attackers, rendering even sophisticated cryptographic systems insecure.

## Multisig and Longer Seed Phrases
It is tempting to assume that, say, a 2-of-3 multisig setup where each share is a 12-word seed is equivalent to the security of a 24-word seed phrase. This is not correct, because the shares are independent.

By way of analogy, guessing seed phrases is like rolling sixes with dice. The effort of guessing two seeds is like re-rolling two dice individually until both show sixes, whereas the effort of guessing a single 24-word seed is like re-rolling two dice simultaneously. As the chance of hitting two sixes at once is significantly harder than hitting a six with one die and another six with a second die, the 24-word seed is more secure.

However, even longer seed phrases don't help; if the randomness source is weak enough or even malicious, guessing the seed phrase requires very little effort regardless of seed length.

## Reliable Sources of Entropy
Two sources can generally be considered reliable for generating entropy for Bitcoin private keys: Bitcoin Core and manual entropy methods.

### Bitcoin Core: Rigorously Audited but not User-Friendly
Bitcoin Core undergoes extensive scrutiny by a large community of skilled developers, making it one of the most audited and trusted Bitcoin implementations, including its randomness generator. However, there are practical challenges in using Bitcoin Core for entropy:

- User Interface Limitations: Bitcoin Core's user interface is not particularly user-friendly, especially for those not well-versed in technical aspects of Bitcoin.
- Air-Gapped Security Requirement: For optimal security, Bitcoin Core should be run on an air-gapped computer that will never connect to the internet again. This requirement can be cumbersome and is not feasible for everyone.

### Manual Entropy Methods: Accessible but Prone to Human Error
Manual entropy methods, such as seed generation through dice or cards, offer a more accessible way to generate entropy. These methods involve physically rolling dice, picking cards, or using a similar randomization method to create a sequence of words or characters. While these methods are user-friendly and don't require specialized software or hardware, they have their own drawbacks:

- Human Bias: Even when trying to be random, humans can introduce bias into the process, which can reduce the entropy of the generated keys.
- Error Potential: Manual methods are susceptible to human error, such as misreading dice or making transcription errors, which can compromise the security of the generated keys.
- Biased Equipment: Not all dice are created equal. Many dice, especially those not meeting casino-grade standards, can have slight biases that affect their roll outcomes. If picking out of a jar, the amount of free air, method of shaking/stirring, wear and tear on the pills/balls, however subtle, can also reduce the entropy.

## Limitations of Entropy in Hardware Wallets
Many users might wonder why not simply rely on the entropy provided by dedicated Bitcoin hardware wallets. The main reason is that hardware wallets are susceptible to supply chain attacks. Because of the nature of randomness, it is practically impossible to detect whether a particular seed is truly random or generated by malicious firmware or hardware installed during production or delivery of the device.

Even if a device contains no backdoors, its random number generator is most likely not nearly as scrutinized as Bitcoin Core's implementation. An accidentally broken generator is as bad as a malicious one.

In conclusion, the strength of a multisig arrangement is heavily dependent on the entropy of the individual keys. Without high-quality entropy, even a multisig setup cannot guarantee robust security against hacks.

*Keep in mind that a multisig wallet consisting of high entropy shares is, of course, more secure than a singlesig with high entropy. And the physical distribution of the shares will also protect against physical theft, which has nothing to do with the entropy.*