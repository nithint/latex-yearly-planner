\begin{tabularx}{\linewidth}{l|X}
  \arrayrulecolor{\myColorGray}
{{ range $todo := .Body.Todos }}
  {{ $todo.HyperLink }} & \myLineHeightButLine{} \\ \hline
{{ end }}
\end{tabularx}