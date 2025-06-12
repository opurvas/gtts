# 🗣️ Text-to-Speech (TTS) Flask API using gTTS

This project provides a simple REST API to convert input text (in various languages) into speech using Google Text-to-Speech (`gTTS`). The output is saved as an MP3 file, which can be downloaded or accessed via a URL.

---

## 🚀 Features

- Accepts JSON input with text and language code
- Converts text to speech using `gTTS`
- Saves output as an MP3 file
- Returns download link to the generated MP3
- Multi-language support (English, Hindi, Spanish, etc.)

---

## 📦 Requirements

Install the required Python packages using:

```python
pip install -r requirements.txt


```
▶️ Running the API

Run the Flask server using:

```python 
python app.py

```
By default, the server runs on:

```python

http://localhost:8080

```



---

###  📥 API Usage

Include endpoint info, request format, and a sample:


### `POST /api/tts`

Send a JSON payload like:

```json
{
  "text": "Hello, this is a test",
  "lang": "en"
}
