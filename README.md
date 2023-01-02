<p style='text-align: justify;'>

## ASCII Art Web

| Column 1 | Column 2 |
| --- | --- |
| ![Ascii Art Web](aaw0.png) | ![Download button](aaw1.png) |


### Description

ASCII Art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.

The current project includes following subprojects: 
-   [x] stylize
-   [x] dockerize
-   [x] exportfile

Folder <code>res</code> contains several functions, files and structures required for the main code.

Folder <code>templates</code> contains html and css files for the server.

### Allowed Packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed.

### Audit Details

- Here you can see [audit details](https://github.com/01-edu/public/tree/master/subjects/ascii-art-web/audit).

### Usage

- To run the code type in your terminal:
```
go run main.go
```
- Open your web browser and search http://localhost:8080
- Enjoy!
- To stop the server press in your terminal:
```
Ctrl + C
```

Text area on the left side shows the generated ASCII art output, while the text area on the right is for user input. Desired banner style can be chosen by selecting a corresponding radio button.

To generate ASCII art press the <code>Generate ASCII Art</code> button with the text typed and the banner selected.

The program implements built into the home page form with POST method attribute which after giving the appropriate values to the text areas and radio buttons sends them to the server through the /ascii-art subpage, performs all the required calculations and redirects back to the home page where the results are taken from the global variable.

You can perform a hard refresh of the localhost web page in case some styles are not loading. For this, press the <code>Ctrl + F5</code>.

### Stylize

Now, you can drag text-input and banner-selection blocks by holding your left mouse button over the block titles. CSS-flexbox approach has been used when writing initial css which means the design should be responsive.

### Export File

- Type in your text and select the banner.
- Press <code>Download</code> button to download the file.txt with you ASCII artwork.
- Press <code>Generate ASCII Art</code> and compare the results.

### Group Members

- fpetuhov (okuu)
- jjelisav
</p>