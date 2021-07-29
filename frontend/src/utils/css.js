export default function getRule(rules) {
  rules = rules.map((rule) => rule.toLowerCase());

  return [...document.styleSheets]
    .flatMap((styleSheet) => {
      try {
        return [...styleSheet.cssRules];
      } catch {
        return [];
      }
    })
    .filter((cssRule) => cssRule instanceof window.CSSStyleRule)
    .find(
      (cssRule) => rules.indexOf(cssRule.selectorText.toLowerCase()) !== -1
    );
}
