from flask import Flask, request, jsonify
from gtts import gTTS
from datetime import datetime
import os

app = Flask(__name__)

# Ensure static folder exists
if not os.path.exists("static"):
    os.makedirs("static")

@app.route('/api/tts', methods=['POST'])
def text_to_speech():
    try:
        data = request.get_json()
        if not data or 'text' not in data:
            return jsonify({'error': 'Text field is required in JSON'}), 400

        text = data['text']
        lang = data.get('lang', 'en')  # default language is English

        # Generate unique filename
        timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
        filename = f"tts_{timestamp}.mp3"
        output_path = os.path.join("static", filename)

        # Generate TTS audio
        tts = gTTS(text=text, lang=lang)
        tts.save(output_path)

        return jsonify({
            'message': 'MP3 file generated successfully',
            'file_path': output_path,
            'file_url': f'/static/{filename}'
        }), 200

    except Exception as e:
        return jsonify({'error': str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True, port=8080)
