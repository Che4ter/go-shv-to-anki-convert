# go-shv-to-anki-convert
Converts SHV exam questions to .csv for easy anki import.

## Usage
1. Compile with ``make build``
2. Go to the SHV E-Learning and select the desired questions (https://www.shv-fsvl.ch/mein-account/e-learning/#/settings)
3. Open Devtools in Firefox, go to the Network Tab and activate "Persist Logs" 
4. Start the Exam
5. Copy the SOAP Response from the request which contains the questions and save as xml
6. Run the converter with ``./bin/shvtoanki -xml=eLearning.xml -category=Fluglehre``
7. The utility converts the .xml into a csv and also downloads all question images into the folder ``images`` use the flag ``-imgPath`` to provide a different folder
8. Copy all images from the image folder to ``~/.local/share/Anki2/User 1/collection.media``
9. In Anki, click Import, 
   - Select the generated csv file
   - Type: Basic
   - Deck: Choose Deck
   - leave the rest
   - Import
10. Have fun learning :)
