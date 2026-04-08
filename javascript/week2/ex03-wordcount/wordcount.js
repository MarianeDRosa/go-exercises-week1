// @ts-nocheck
function countWords(text) {
  const counts = {};
  // Divide por espaços e remove vazios, transformando tudo em minúsculo
  const words = text.toLowerCase().split(/\s+/).filter(w => w !== "");

  words.forEach(word => {
    // Se a palavra já existe, soma 1. Se não, começa com 1.
    counts[word] = (counts[word] || 0) + 1;
  });

  return counts;
}

module.exports = { countWords };