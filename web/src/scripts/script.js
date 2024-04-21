document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('videoInput').addEventListener('change', function() {
        var input = this;
        if (input.files && input.files[0]) {
            var reader = new FileReader();
            reader.onload = function(e) {
                document.querySelector('.image-upload-wrap').style.display = 'none';
                document.querySelector('.file-upload-image').setAttribute('src', e.target.result);
                document.querySelector('.file-upload-content').style.display = 'block';
                document.querySelector('.image-title').textContent = input.files[0].name;
            };
            reader.readAsDataURL(input.files[0]);
        } else {
            removeUpload();
        }
    });

    document.getElementById('removeBtn').addEventListener('click', function() {
        removeUpload();
    });

    function removeUpload() {
        var uploadWrap = document.querySelector('.image-upload-wrap');
        var uploadContent = document.querySelector('.file-upload-content');
        uploadWrap.style.display = 'block';
        uploadContent.style.display = 'none';

        var dragWrap = document.querySelector('.image-upload-wrap');
        dragWrap.addEventListener('dragover', function() {
            dragWrap.classList.add('image-dropping');
        });

        dragWrap.addEventListener('dragleave', function() {
            dragWrap.classList.remove('image-dropping');
        });
    }
});

document.getElementById('videoForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    const file = document.getElementById('videoInput').files[0];
    const formData = new FormData();
    formData.append('video', file);
    await fetch('/api/upload-video', {
        method: 'POST',
        body: formData
    });
    alert('Video uploaded!');
});
