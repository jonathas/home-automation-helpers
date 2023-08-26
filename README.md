# Home Automation Helpers

This is a REST API which contains endpoints I've implemented to help me with my home automations.

It's intended to run locally and be called locally.

## Docs

A Postman collection can be found inside of the /docs directory

## Configuration

- Set your VoiceMonkey token in the .env file
- Compile the project and keep the .env file next to it when running
- A systemd service file is provided inside of the /resources directory
- Copy it to /etc/systemd/system
- Run:
  - sudo systemctl daemon-reload
  - sudo systemctl enable home-automation-helpers
  - sudo systemctl start home-automation-helpers
