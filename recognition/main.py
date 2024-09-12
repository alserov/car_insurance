import os
import random
from werkzeug.utils import secure_filename as sf
from flask import Flask, request, jsonify

app = Flask(__name__)

# create uploads folder if not exists
UPLOAD_FOLDER = 'uploads'
if not os.path.exists(UPLOAD_FOLDER):
    os.makedirs(UPLOAD_FOLDER)
app.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER

# route for image recognition
@app.route('/check', methods=['POST'])
def checkIfIsAbleToBeInsured():
    if 'image' not in request.files:
        return jsonify({'error': 'image is not provided'}), 400

    file = request.files['image']

    filename = sf(file.filename)
    file_path = os.path.join(app.config['UPLOAD_FOLDER'], filename)
    file.save(file_path)

    mult = random.uniform(0, 1)

    os.remove(file_path)

    if mult > 0.5:
        return jsonify({'insurable': True})

    return jsonify({'insurable': False})

@app.route('/detect_multiplier', methods=['POST'])
def detectDamageMultiplier():
    if 'image' not in request.files:
        return jsonify({'error': 'image is not provided'}), 400

    file = request.files['image']

    filename = sf(file.filename)
    file_path = os.path.join(app.config['UPLOAD_FOLDER'], filename)
    file.save(file_path)

    os.remove(file_path)

    return jsonify({'score': random.uniform(1.5, 1.9)})


# starting server
if __name__ == '__main__':
    app.run(debug=True)