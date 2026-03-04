# Gator: CLI RSS AggreGATOR

## Requirements
- Install recent version of GOlang
- Install recent version of Postgres

## Installation
- To install gator use command: go install gator

## Set Up
- Create .gatorconfig.json in your home directory
- Update the configuration file DBUrl field with your Postgres connection string
- Register your initial user with command: gator register <username>

## Usage
- Example commands that can be run include:
    - login: to login with registered user and set the  user as active in the session
    - users: lists all users currently registered
    - addFeed: adds RSS feed using 2 arguments (name and url) to current user and follows
    - follow: attaches user to follow feed 
    - following: lists all feeds the user is following
    - unfollow: removes user from feed follow
    - agg: scrapes RSS feeds and saves posts
    - browse: browse posts saved (accepts number of posts to browse)