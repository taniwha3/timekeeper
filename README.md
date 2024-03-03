# WARNING
This project is under active development without a v1 release. Expect the data structures and storage to change.

# TimeKeeper
Welcome to TimeKeeper, a simple CLI application designed to help you manage and track your project activities with ease. TimeKeeper allows you to add projects to a list, select them through a menu, and log your activities with timestamps. It's a perfect tool for freelancers, developers, or anyone looking to keep a detailed record of their project engagements.

## Features
* Project Management: Easily add your projects to a projects.txt file to keep them organized.
* Interactive Menu: Upon running the application, your projects will be displayed in a user-friendly menu for easy selection.
* Activity Logging: TimeKeeper logs when you start working on a project, switch between projects, and when you stop working on a project.
* Continuous Tracking: The application logs your current project selection or the lack of selection every minute, ensuring detailed tracking of your time.
* Data Persistence: Each selection is recorded in a unique file within the `data` directory, providing a clear history of your project activities.
* Auto-Deselection: To prevent accidental logging, TimeKeeper automatically deselects the current project if your computer goes to sleep and comes back online after 5 and a half minutes.

## Data Structure
Each activity log is stored in a structured format containing the projects list, the currently selected project, and the timestamp of the log. Here's an example of the data structure:

```json
{
  "selected": "one",
  "timestamp": "2024-03-03T12:06:52.946635857-08:00"
}
```

## Getting Started
To get started with TimeKeeper, follow these steps:

* Installation: Clone this repository to your local machine.
* Add Projects: Open the projects.txt file and add the names of the projects you want to track. Each project should be on a new line.
* Run TimeKeeper: Execute the application to see your projects listed in the menu. Use the menu to select or deselect projects as you work on them.
* Review Logs: Navigate to the `data` directory to view the unique files generated for each of your selections, containing detailed logs of your project activities.

## Contributions
We welcome contributions to TimeKeeper! If you have suggestions for improvements or encounter any issues, please feel free to open an issue or submit a pull request.

## License
TimeKeeper is released under the ApacheV2 License. See the LICENSE file for more details.
