document.querySelector('input[type=file]').addEventListener('change', function() {
    const file = this.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onloadend = function() {
            document.getElementById('preview').src = reader.result;
        }
        reader.readAsDataURL(file);
    } else {
        document.getElementById('preview').src = "";
    }
});

function recognizeImage() {
    const file = document.querySelector('input[type=file]').files[0];
    if (!file) {
        alert('请先选择一个文件。');
        return;
    }

    const reader = new FileReader();
    reader.onload = function() {
        const base64Image = reader.result.replace(/^data:image\/(.*);base64,/, '');
        sendImageForClassification(base64Image);
    };
    reader.readAsDataURL(file);
}

function sendImageForClassification(base64Image) {
    const startTime = performance.now();

    const xhr = new XMLHttpRequest();
    xhr.open('POST', '/api/image/classify', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onload = function() {
        if (xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            const results = response.results;
            displayResults(results, startTime);
        } else {
            alert('发生错误，无法识别图像。');
        }
    };
    xhr.onerror = function() {
        alert('请求失败，请检查网络连接。');
    };
    xhr.send(JSON.stringify({ image: base64Image }));
}

function displayResults(results, startTime) {
    const resultDiv = document.getElementById('result');
    const responseTimeDiv = document.getElementById('response-time');
    const endTime = performance.now();
    const responseTime = endTime - startTime;

    if (!Array.isArray(results)) {
        resultDiv.innerHTML = "未能识别到结果或响应格式不正确。";
        return; // 如果results不是数组，则直接返回
    }

    responseTimeDiv.textContent = `本次响应用时：${responseTime.toFixed(2)}ms`;
    resultDiv.innerHTML = results.map((result, index) => {
        const colorStyle = index === 0 ? 'color:red;' : '';
        return `<span style="${colorStyle}">识别结果：${result.name}（置信度：${result.score}）</span><br>`;
    }).join('');
}