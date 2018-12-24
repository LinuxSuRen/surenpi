var Clipboard = require('clipboard/dist/clipboard.js');
new Clipboard('.copy', {
  target: function(trigger) {
    return trigger.nextElementSibling;
  }
  }).on('success', function(e) {
    successMessage(e.trigger, 'Copied!');
    e.clearSelection();
  }).on('error', function(e) {
    successMessage(e.trigger, fallbackMessage(e.action));
});

function successMessage(elem, msg) {
  elem.setAttribute('class', 'copied bg-primary-color-dark f6 top-0 right-0 lh-solid hover-bg-primary-color-dark bn white ph3 pv2');
  elem.setAttribute('aria-label', msg);
  window.setTimeout(function(){
    var clazz = elem.getAttribute('class');
    clazz = clazz.replace('copied', 'copy');
    elem.setAttribute('class', clazz);
  }, 3000);
}

function fallbackMessage(elem, action) {
  var actionMsg = '';
  var actionKey = (action === 'cut' ? 'X' : 'C');
  if (isMac) {
      actionMsg = 'Press ⌘-' + actionKey;
  } else {
      actionMsg = 'Press Ctrl-' + actionKey;
  }
  return actionMsg;
}