**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-22 13:53:00 UTC（2026-03-22 21:53:00 UTC+8）

**代理总数：67**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 669ms | 否 | ✓ 943ms | ✓ 889ms | ✓ 883ms | http |
| 113.160.132.26:8080 | ✓ 1814ms | 否 | ✓ 1264ms | ✓ 1561ms | ✓ 1100ms | http |
| 45.167.124.52:8080 | ✓ 787ms | ✓ 1797ms | ✓ 1390ms | 否 | ✓ 1482ms | http |
| 85.198.96.242:3128 | ✓ 1060ms | 否 | ✓ 1252ms | 否 | ✓ 1380ms | http |
| 167.103.34.108:8800 | ✓ 1138ms | 否 | ✓ 1190ms | ✓ 1357ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1901ms | 否 | ✓ 1798ms | ✓ 1919ms | ✓ 1990ms | http |
| 137.220.151.110:6005 | ✓ 1667ms | 否 | ✓ 1068ms | ✓ 1481ms | ✓ 1265ms | http |
| 120.92.212.16:8890 | ✓ 1191ms | 否 | ✓ 956ms | ✓ 1264ms | 否 | http |
| 37.187.109.70:10111 | ✓ 1110ms | 否 | ✓ 1626ms | 否 | ✓ 1990ms | http |
| 106.75.15.167:7890 | ✓ 1143ms | ✓ 1207ms | ✓ 901ms | 否 | 否 | http |
| 45.136.131.51:8449 | 否 | ✓ 969ms | ✓ 926ms | ✓ 769ms | ✓ 619ms | http |
| 147.161.239.240:8800 | ✓ 1085ms | ✓ 1751ms | ✓ 1534ms | 否 | 否 | http |
| 101.43.127.100:8877 | ✓ 903ms | ✓ 1085ms | ✓ 1403ms | 否 | ✓ 934ms | http |
| 106.14.203.63:3333 | ✓ 1522ms | ✓ 1983ms | ✓ 1798ms | ✓ 1912ms | ✓ 896ms | http |
| 34.150.20.6:8888 | ✓ 703ms | ✓ 1055ms | ✓ 751ms | ✓ 877ms | ✓ 707ms | http |
| 45.129.141.143:3128 | ✓ 1119ms | 否 | 否 | ✓ 1959ms | ✓ 1768ms | http |
| 120.92.212.16:7890 | ✓ 942ms | ✓ 1448ms | ✓ 967ms | ✓ 1261ms | ✓ 1207ms | http |
| 64.227.76.27:1080 | ✓ 778ms | 否 | ✓ 1708ms | ✓ 1773ms | 否 | http |
| 45.136.131.35:8452 | 否 | ✓ 1206ms | ✓ 1065ms | ✓ 903ms | ✓ 648ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1151ms | ✓ 1277ms | ✓ 1276ms | http |
| 142.171.224.229:7890 | ✓ 1548ms | ✓ 867ms | ✓ 504ms | ✓ 705ms | ✓ 544ms | http |
| 38.145.208.191:8453 | ✓ 1421ms | ✓ 753ms | ✓ 1106ms | ✓ 1894ms | ✓ 641ms | http |
| 45.136.131.32:8452 | ✓ 1420ms | 否 | ✓ 618ms | ✓ 1279ms | ✓ 1826ms | http |
| 5.102.109.41:999 | ✓ 1957ms | 否 | ✓ 972ms | 否 | ✓ 1114ms | http |
| 20.210.39.153:8561 | ✓ 1112ms | ✓ 1277ms | ✓ 489ms | ✓ 936ms | ✓ 717ms | http |
| 20.78.26.206:8561 | ✓ 1110ms | 否 | ✓ 492ms | ✓ 794ms | ✓ 626ms | http |
| 77.232.135.22:1080 | ✓ 864ms | 否 | ✓ 1937ms | ✓ 1758ms | 否 | http |
| 35.225.22.61:80 | ✓ 1057ms | 否 | ✓ 1233ms | 否 | ✓ 885ms | http |
| 20.27.13.35:8561 | ✓ 1235ms | ✓ 1241ms | ✓ 540ms | ✓ 850ms | ✓ 690ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1809ms | ✓ 1820ms | ✓ 1182ms | ✓ 1716ms | http |
| 20.27.14.220:8561 | ✓ 494ms | ✓ 1427ms | ✓ 491ms | ✓ 816ms | ✓ 664ms | http |
| 20.27.11.248:8561 | ✓ 1423ms | ✓ 1350ms | ✓ 654ms | ✓ 857ms | ✓ 626ms | http |
| 166.88.55.83:7890 | ✓ 653ms | ✓ 1108ms | ✓ 658ms | ✓ 825ms | ✓ 672ms | http |
| 106.117.208.101:7890 | ✓ 1990ms | ✓ 1578ms | ✓ 1466ms | 否 | ✓ 995ms | http |
| 103.113.70.189:1081 | ✓ 368ms | ✓ 1775ms | ✓ 1539ms | 否 | ✓ 891ms | http |
| 47.77.193.180:1080 | ✓ 140ms | 否 | ✓ 277ms | ✓ 741ms | ✓ 549ms | http |
| 16.78.119.130:443 | 否 | 否 | ✓ 1850ms | ✓ 1867ms | ✓ 1616ms | http |
| 218.89.134.230:3333 | ✓ 1528ms | 否 | ✓ 1855ms | 否 | ✓ 1431ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1146ms | ✓ 1622ms | ✓ 1037ms | ✓ 975ms | http |
| 59.8.203.55:80 | 否 | 否 | ✓ 1088ms | ✓ 1008ms | ✓ 789ms | http |
| 149.62.191.202:3128 | ✓ 1186ms | 否 | ✓ 1372ms | ✓ 1948ms | ✓ 1466ms | http |
| 160.250.4.245:1 | ✓ 1740ms | 否 | ✓ 1300ms | ✓ 1394ms | ✓ 1188ms | http |
| 207.254.71.62:8088 | ✓ 1359ms | 否 | ✓ 1687ms | 否 | ✓ 1878ms | http |
| 62.234.206.73:3128 | ✓ 1438ms | ✓ 1248ms | 否 | ✓ 1199ms | 否 | http |
| 172.212.68.37:3128 | ✓ 1241ms | ✓ 1824ms | ✓ 677ms | ✓ 1489ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1727ms | 否 | 否 | ✓ 1795ms | ✓ 1707ms | http |
| 45.151.183.183:1080 | ✓ 871ms | 否 | ✓ 1323ms | 否 | ✓ 1418ms | http |
| 45.144.28.81:10808 | ✓ 940ms | 否 | ✓ 1732ms | 否 | ✓ 1160ms | http |
| 137.220.150.22:6005 | ✓ 1664ms | 否 | 否 | ✓ 1392ms | ✓ 1015ms | http |
| 218.60.0.214:80 | 否 | ✓ 1749ms | ✓ 1417ms | ✓ 1499ms | ✓ 1080ms | http |
| 103.69.84.106:3131 | ✓ 1829ms | 否 | ✓ 1251ms | ✓ 1237ms | ✓ 917ms | http |
| 104.248.81.109:3128 | ✓ 633ms | 否 | ✓ 1930ms | ✓ 1397ms | ✓ 1141ms | http |
| 20.27.15.111:8561 | ✓ 491ms | ✓ 1197ms | ✓ 486ms | ✓ 840ms | ✓ 637ms | http |
| 202.35.251.72:8080 | ✓ 1671ms | 否 | ✓ 1934ms | ✓ 1279ms | ✓ 981ms | http |
| 54.37.72.89:80 | ✓ 716ms | 否 | ✓ 1783ms | 否 | ✓ 1645ms | http |
| 114.237.77.231:1080 | ✓ 1691ms | ✓ 1246ms | ✓ 1032ms | ✓ 1930ms | 否 | http |
| 207.244.244.178:3128 | ✓ 377ms | 否 | ✓ 1772ms | ✓ 1122ms | ✓ 985ms | http |
| 45.149.92.147:5001 | ✓ 654ms | 否 | ✓ 660ms | ✓ 826ms | ✓ 672ms | http |
| 116.80.65.85:3172 | 否 | 否 | ✓ 1526ms | ✓ 1875ms | ✓ 1665ms | http |
| 116.80.65.79:3172 | ✓ 1543ms | 否 | ✓ 1565ms | 否 | ✓ 1801ms | http |
| 139.159.99.242:8080 | ✓ 819ms | ✓ 1040ms | ✓ 844ms | 否 | 否 | http |
| 103.135.102.161:8080 | ✓ 1686ms | 否 | ✓ 1941ms | ✓ 1296ms | ✓ 975ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1564ms | ✓ 1637ms | ✓ 1304ms | 否 | http |
| 193.23.200.251:10808 | ✓ 900ms | ✓ 1981ms | ✓ 1315ms | 否 | ✓ 1761ms | http |
| 137.184.1.155:3128 | ✓ 323ms | ✓ 1460ms | ✓ 317ms | ✓ 747ms | ✓ 589ms | http |
| 107.178.115.140:3128 | ✓ 495ms | 否 | ✓ 850ms | ✓ 1373ms | ✓ 881ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 1342ms | ✓ 866ms | ✓ 718ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
