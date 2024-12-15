---
title: Is it time to code your config?
description: My article description
tags: ''
cover_image: ''
canonical_url: null
published: false
id: 2157929
---

Intro: for most people, simple config file with ini, yaml or other is way enough for their usage.

12 factor app: important that you don't hardcode everything in your code, you have configuration coming from env

Once you start to be in projects with more scale and way more different configuration layout (like different datacenters, regions, zones, setups & so on), configuration is way more complex to manage and, as we will see, standard tools lacks of features to manage this complexity.

This article introduces a reflection on this subject, statuting issue.

# What does static configuration format lack?

When we mention "classic configuration" here, we are talking about  JSON/YAML/TOML like
We will not mention Env variable and program arguments since they contains less possibilities compared to what config file can perform

Here we will talk mostly about YAML since it's a very popular language for configuration these days ()

## A proper language to manage basic coding

Go to helm Templating like Go Template or Jinja!
* Don't ensure easily that the exporter file is a correct YAML/JSON/
* Your code starts to be a melt of YAML and templating language, hard for human to understanding what's going on exactly;
* Testing... ah ah ah

A descriptive language like YAML is not 

## Validate your config

## Hierarchical configuration

You have YAML config for DRY situation, but it's scoped to a single file.
For multiple file, you can have custom stuff (like GitLab include) but it's always custom and lack of 

## Configuration from other locations

## IDE integration

## When config comes dynamically from env

* Secrets
* URLs of services
* Database & broker location

## Why not directly a programming language?

* Security

# What an ideal config tool could be

# Some interesting ideas

Quick overview and rapid tests on following elements:

* TODO 

## Pkl

Cons: you can do too much stuff (whole generator are coded in Pkl). I'm afraid it can tends to be like templates

## Cue

## Kcl

# So, should you move to "Config as Code"?

Like always, it depends on your need. But it you have complex deployment structure, with a lot of possilibities, my opinion is good to have a try. I haven't tried that deep for 
