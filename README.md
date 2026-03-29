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

最后更新：2026-03-29 14:02:13 UTC（2026-03-29 22:02:13 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 677ms | ✓ 1005ms | ✓ 775ms | ✓ 855ms | ✓ 677ms | http |
| 147.161.210.140:8800 | 否 | 否 | ✓ 778ms | ✓ 1482ms | ✓ 1039ms | http |
| 167.103.34.108:8800 | ✓ 1501ms | 否 | ✓ 1533ms | ✓ 1384ms | 否 | http |
| 147.161.239.240:8800 | ✓ 1097ms | 否 | 否 | ✓ 1651ms | ✓ 1554ms | http |
| 95.213.217.168:52004 | ✓ 1126ms | 否 | ✓ 1631ms | ✓ 1847ms | ✓ 1499ms | http |
| 167.172.77.49:8080 | ✓ 1560ms | 否 | ✓ 1539ms | ✓ 1143ms | 否 | http |
| 167.103.115.102:8800 | ✓ 1571ms | 否 | 否 | ✓ 1794ms | ✓ 1015ms | http |
| 42.96.16.158:1311 | ✓ 1679ms | 否 | 否 | ✓ 1297ms | ✓ 1236ms | http |
| 35.225.22.61:80 | ✓ 416ms | 否 | ✓ 969ms | ✓ 1060ms | 否 | http |
| 39.185.46.193:5911 | ✓ 1739ms | ✓ 862ms | ✓ 942ms | ✓ 926ms | ✓ 724ms | http |
| 113.160.132.26:8080 | ✓ 901ms | 否 | ✓ 928ms | ✓ 1282ms | 否 | http |
| 1.231.81.166:3128 | ✓ 864ms | 否 | ✓ 1303ms | ✓ 1595ms | ✓ 839ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 1055ms | ✓ 945ms | ✓ 880ms | http |
| 167.103.144.127:8800 | ✓ 1729ms | 否 | ✓ 1273ms | 否 | ✓ 1535ms | http |
| 45.12.151.226:2829 | ✓ 796ms | 否 | ✓ 1308ms | 否 | ✓ 1836ms | http |
| 167.103.31.122:8800 | ✓ 1609ms | 否 | ✓ 1367ms | 否 | ✓ 1550ms | http |
| 160.238.65.6:3128 | ✓ 1037ms | 否 | ✓ 1276ms | ✓ 1790ms | ✓ 1096ms | http |
| 160.238.65.9:3128 | ✓ 1023ms | ✓ 1947ms | ✓ 1328ms | ✓ 1782ms | 否 | http |
| 103.84.95.54:7890 | ✓ 790ms | 否 | ✓ 915ms | ✓ 1237ms | ✓ 887ms | http |
| 106.75.15.167:7890 | ✓ 1142ms | ✓ 1247ms | 否 | 否 | ✓ 1866ms | http |
| 103.113.70.189:1081 | ✓ 467ms | ✓ 1282ms | ✓ 810ms | 否 | ✓ 1035ms | http |
| 120.92.212.16:7890 | ✓ 1215ms | ✓ 1458ms | 否 | ✓ 1294ms | ✓ 999ms | http |
| 101.43.127.100:8877 | ✓ 1590ms | 否 | 否 | ✓ 1240ms | ✓ 1628ms | http |
| 177.234.217.88:999 | 否 | 否 | ✓ 1770ms | ✓ 1962ms | ✓ 1709ms | http |
| 38.145.208.216:8448 | ✓ 1218ms | ✓ 1802ms | 否 | ✓ 1841ms | ✓ 1534ms | http |
| 20.27.14.220:8561 | 否 | ✓ 991ms | ✓ 656ms | ✓ 882ms | ✓ 718ms | http |
| 20.210.76.175:8561 | 否 | 否 | ✓ 512ms | ✓ 886ms | ✓ 673ms | http |
| 45.136.198.40:3128 | ✓ 1159ms | ✓ 1827ms | ✓ 1804ms | 否 | ✓ 1726ms | http |
| 34.101.184.164:3128 | ✓ 1751ms | 否 | ✓ 1384ms | ✓ 1820ms | ✓ 1242ms | http |
| 38.34.179.91:8444 | ✓ 248ms | 否 | ✓ 1813ms | ✓ 968ms | ✓ 903ms | http |
| 210.223.44.230:3128 | ✓ 596ms | ✓ 816ms | 否 | ✓ 972ms | 否 | http |
| 20.27.15.49:8561 | 否 | ✓ 1145ms | ✓ 783ms | ✓ 869ms | ✓ 953ms | http |
| 45.136.130.250:8450 | 否 | ✓ 1561ms | ✓ 410ms | 否 | ✓ 607ms | http |
| 20.78.26.206:8561 | ✓ 497ms | ✓ 848ms | ✓ 491ms | ✓ 904ms | ✓ 768ms | http |
| 20.210.39.153:8561 | ✓ 494ms | ✓ 1429ms | ✓ 563ms | ✓ 798ms | ✓ 666ms | http |
| 20.27.15.111:8561 | ✓ 615ms | ✓ 1051ms | ✓ 845ms | ✓ 1011ms | ✓ 776ms | http |
| 20.27.11.248:8561 | ✓ 617ms | ✓ 1052ms | ✓ 831ms | ✓ 1018ms | ✓ 781ms | http |
| 20.27.13.35:8561 | ✓ 607ms | 否 | ✓ 615ms | ✓ 972ms | ✓ 808ms | http |
| 38.34.179.172:8451 | 否 | 否 | ✓ 179ms | ✓ 783ms | ✓ 611ms | http |
| 217.76.245.80:999 | ✓ 839ms | 否 | ✓ 1241ms | ✓ 1558ms | 否 | http |
| 139.135.81.163:8082 | ✓ 1424ms | 否 | ✓ 1974ms | ✓ 1858ms | ✓ 1645ms | http |
| 101.32.244.83:8080 | ✓ 1482ms | 否 | ✓ 950ms | ✓ 1449ms | ✓ 1241ms | http |
| 121.43.196.213:8222 | ✓ 947ms | ✓ 1085ms | ✓ 910ms | ✓ 1183ms | ✓ 895ms | http |
| 121.43.196.210:8222 | ✓ 962ms | ✓ 1074ms | ✓ 902ms | ✓ 1201ms | ✓ 925ms | http |
| 114.55.226.123:10086 | ✓ 1023ms | ✓ 1834ms | ✓ 1054ms | ✓ 1260ms | ✓ 1080ms | http |
| 160.238.65.4:3128 | ✓ 1552ms | ✓ 1884ms | ✓ 1815ms | 否 | ✓ 1761ms | http |
| 160.238.65.5:3128 | ✓ 1553ms | ✓ 1860ms | ✓ 1848ms | 否 | ✓ 1729ms | http |
| 38.145.208.178:8452 | ✓ 588ms | 否 | ✓ 340ms | ✓ 813ms | ✓ 740ms | http |
| 38.145.208.224:8446 | ✓ 595ms | 否 | ✓ 830ms | ✓ 995ms | ✓ 688ms | http |
| 38.145.203.132:8450 | ✓ 162ms | ✓ 717ms | ✓ 456ms | 否 | 否 | http |
| 45.136.131.47:8445 | ✓ 409ms | ✓ 1356ms | ✓ 1396ms | 否 | ✓ 879ms | http |
| 38.145.208.204:8451 | ✓ 345ms | 否 | ✓ 1082ms | ✓ 796ms | ✓ 746ms | http |
| 101.47.73.135:3128 | ✓ 1076ms | 否 | ✓ 1647ms | ✓ 1904ms | ✓ 1212ms | http |
| 208.87.243.199:7878 | ✓ 1524ms | 否 | ✓ 981ms | 否 | ✓ 1309ms | http |
| 45.140.147.82:1081 | ✓ 612ms | ✓ 1482ms | ✓ 1670ms | ✓ 1852ms | ✓ 1209ms | http |
| 45.149.92.147:5001 | 否 | 否 | ✓ 1847ms | ✓ 1518ms | ✓ 776ms | http |
| 106.117.208.101:7890 | ✓ 1728ms | 否 | ✓ 1043ms | 否 | ✓ 1209ms | http |
| 120.92.212.16:8890 | ✓ 1179ms | ✓ 1248ms | ✓ 960ms | ✓ 1225ms | ✓ 1203ms | http |
| 45.140.147.82:1082 | ✓ 554ms | 否 | ✓ 1497ms | ✓ 1708ms | ✓ 1136ms | http |
| 115.231.181.40:8128 | ✓ 1845ms | ✓ 1885ms | 否 | ✓ 1851ms | ✓ 1027ms | http |
| 8.219.97.248:80 | ✓ 1395ms | 否 | ✓ 1795ms | ✓ 1277ms | 否 | http |
| 103.82.23.118:5253 | 否 | 否 | ✓ 1380ms | ✓ 1730ms | ✓ 1431ms | http |
| 64.227.76.27:1080 | ✓ 1506ms | 否 | ✓ 1499ms | 否 | ✓ 1895ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 853ms | ✓ 1153ms | ✓ 931ms | http |
| 38.34.179.105:8451 | ✓ 1895ms | ✓ 1050ms | ✓ 970ms | 否 | ✓ 666ms | http |
| 20.78.118.91:8561 | 否 | ✓ 1733ms | ✓ 823ms | ✓ 1265ms | ✓ 1154ms | http |
| 193.233.22.29:10808 | ✓ 506ms | 否 | ✓ 622ms | 否 | ✓ 1233ms | http |
| 38.34.179.85:8444 | ✓ 209ms | ✓ 1506ms | ✓ 512ms | ✓ 1017ms | ✓ 734ms | http |
| 38.34.179.87:8444 | ✓ 240ms | 否 | ✓ 429ms | ✓ 768ms | ✓ 605ms | http |
| 38.34.179.89:8444 | ✓ 217ms | 否 | ✓ 426ms | ✓ 1195ms | ✓ 609ms | http |
| 104.248.81.109:3128 | ✓ 559ms | 否 | ✓ 1549ms | ✓ 1494ms | ✓ 1089ms | http |
| 38.34.183.164:8444 | ✓ 244ms | ✓ 739ms | ✓ 410ms | ✓ 1351ms | ✓ 1451ms | http |
| 38.34.179.61:8445 | ✓ 504ms | ✓ 1413ms | ✓ 216ms | ✓ 791ms | ✓ 986ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1256ms | ✓ 1941ms | ✓ 1675ms | http |
| 38.145.218.87:8445 | ✓ 231ms | ✓ 1857ms | ✓ 681ms | ✓ 899ms | ✓ 1291ms | http |
| 38.34.183.211:8444 | ✓ 217ms | ✓ 1693ms | ✓ 1110ms | ✓ 1354ms | ✓ 912ms | http |
| 38.145.203.86:8452 | ✓ 426ms | ✓ 1141ms | ✓ 308ms | 否 | ✓ 1827ms | http |
| 38.145.208.179:8452 | ✓ 789ms | ✓ 1108ms | ✓ 842ms | ✓ 1017ms | ✓ 1590ms | http |
| 175.194.173.105:3128 | 否 | 否 | ✓ 1124ms | ✓ 1010ms | ✓ 776ms | http |
| 139.99.238.95:8080 | ✓ 1534ms | 否 | ✓ 1169ms | ✓ 1824ms | ✓ 1245ms | http |
| 91.233.223.147:3128 | ✓ 1294ms | 否 | ✓ 1259ms | 否 | ✓ 1962ms | http |
| 5.102.109.41:999 | 否 | 否 | ✓ 426ms | ✓ 1210ms | ✓ 1118ms | http |
| 91.238.123.230:8000 | ✓ 513ms | 否 | ✓ 1735ms | ✓ 1785ms | ✓ 1116ms | http |
| 45.136.131.59:8450 | ✓ 1328ms | ✓ 1438ms | ✓ 1101ms | ✓ 1224ms | ✓ 728ms | http |
| 103.82.23.118:5247 | ✓ 1568ms | 否 | 否 | ✓ 1903ms | ✓ 1640ms | http |
| 65.108.203.35:28080 | ✓ 1476ms | 否 | ✓ 1398ms | 否 | ✓ 1870ms | http |

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
