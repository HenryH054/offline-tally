# offline-tally

An offline reconcilable tally tracker.

This will consist of 2 large parts and an optional 3rd.

1) Server
  * A rest API server
  * Allows for multiple simultaneous inputs
  * Allows for recording of offline made edits
  * Reconciles edits that were made offline
  * Basic authentication
2) App
  * A homescreen widget to display the list
  * The ability to set the webpage from in app
  * The ability to set and save credentials
  * Displays a list with the tally name, current total, editing(set and adjust)
3) Web Page
  * Optionally a place to view the tallies from a website
