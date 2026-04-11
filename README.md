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

最后更新：2026-04-11 10:42:56 UTC（2026-04-11 18:42:56 UTC+8）

**代理总数：89**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 680ms | 否 | ✓ 971ms | ✓ 809ms | ✓ 746ms | http |
| 147.161.239.240:8800 | ✓ 739ms | ✓ 1810ms | ✓ 1498ms | ✓ 1594ms | ✓ 1438ms | http |
| 113.160.132.26:8080 | ✓ 1695ms | ✓ 1317ms | ✓ 1089ms | ✓ 1598ms | ✓ 996ms | http |
| 167.103.115.102:8800 | ✓ 1589ms | 否 | ✓ 1154ms | ✓ 1558ms | ✓ 1140ms | http |
| 202.141.161.53:10808 | ✓ 1009ms | ✓ 1192ms | ✓ 1643ms | 否 | 否 | http |
| 167.103.34.108:8800 | ✓ 1497ms | 否 | ✓ 1989ms | ✓ 1712ms | ✓ 1602ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1997ms | 否 | ✓ 1271ms | ✓ 1683ms | http |
| 45.167.124.52:8080 | ✓ 1818ms | 否 | ✓ 1575ms | ✓ 1853ms | 否 | http |
| 35.225.22.61:80 | ✓ 384ms | 否 | ✓ 1674ms | ✓ 1323ms | 否 | http |
| 45.136.130.181:8445 | ✓ 293ms | ✓ 946ms | ✓ 1222ms | ✓ 803ms | ✓ 541ms | http |
| 155.117.18.36:25388 | ✓ 683ms | ✓ 1175ms | ✓ 1404ms | ✓ 1535ms | ✓ 1211ms | http |
| 38.34.179.178:8444 | ✓ 378ms | 否 | ✓ 368ms | ✓ 689ms | ✓ 577ms | http |
| 38.34.179.199:8452 | ✓ 758ms | ✓ 1605ms | ✓ 1734ms | ✓ 1117ms | ✓ 1354ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1152ms | ✓ 1242ms | ✓ 1033ms | http |
| 167.103.144.127:8800 | ✓ 1299ms | 否 | ✓ 734ms | ✓ 1192ms | ✓ 1067ms | http |
| 167.103.31.122:8800 | ✓ 1572ms | 否 | ✓ 1344ms | ✓ 1702ms | 否 | http |
| 91.238.105.43:2023 | ✓ 988ms | ✓ 1815ms | ✓ 1768ms | 否 | 否 | http |
| 104.129.203.245:10026 | ✓ 1216ms | ✓ 685ms | ✓ 538ms | ✓ 749ms | ✓ 564ms | http |
| 104.129.203.245:10139 | ✓ 1216ms | ✓ 685ms | ✓ 1540ms | ✓ 660ms | 否 | http |
| 104.129.203.245:10733 | ✓ 1217ms | ✓ 699ms | ✓ 1571ms | ✓ 705ms | 否 | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1601ms | ✓ 1377ms | ✓ 1200ms | http |
| 218.108.131.186:17890 | ✓ 851ms | ✓ 983ms | ✓ 1340ms | 否 | 否 | http |
| 38.34.179.98:8451 | ✓ 242ms | ✓ 791ms | ✓ 306ms | 否 | 否 | http |
| 45.136.130.192:8450 | ✓ 885ms | ✓ 1746ms | ✓ 733ms | ✓ 766ms | ✓ 933ms | http |
| 101.43.127.100:8877 | ✓ 934ms | ✓ 1005ms | 否 | 否 | ✓ 895ms | http |
| 120.92.212.16:8890 | ✓ 933ms | 否 | ✓ 970ms | 否 | ✓ 1663ms | http |
| 5.196.101.18:3128 | ✓ 1288ms | 否 | ✓ 982ms | 否 | ✓ 1531ms | http |
| 185.76.240.254:10001 | ✓ 1296ms | 否 | ✓ 1328ms | 否 | ✓ 1818ms | http |
| 185.76.240.167:10001 | ✓ 1296ms | 否 | ✓ 1707ms | ✓ 1820ms | 否 | http |
| 45.167.125.21:999 | ✓ 1238ms | 否 | 否 | ✓ 1818ms | ✓ 1818ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1992ms | ✓ 1055ms | ✓ 1576ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1809ms | 否 | ✓ 985ms | ✓ 1823ms | 否 | http |
| 38.145.208.229:8453 | 否 | ✓ 1908ms | ✓ 189ms | ✓ 815ms | ✓ 1049ms | http |
| 1.231.81.166:3128 | ✓ 1383ms | ✓ 1761ms | ✓ 1661ms | ✓ 1842ms | ✓ 1847ms | http |
| 54.222.174.194:80 | 否 | ✓ 1687ms | ✓ 1805ms | ✓ 1562ms | 否 | http |
| 24.144.86.173:1080 | ✓ 766ms | ✓ 962ms | ✓ 256ms | ✓ 727ms | ✓ 547ms | http |
| 45.136.131.64:8445 | ✓ 293ms | ✓ 793ms | ✓ 120ms | ✓ 685ms | ✓ 756ms | http |
| 45.136.131.58:8445 | ✓ 302ms | ✓ 830ms | ✓ 127ms | ✓ 677ms | ✓ 766ms | http |
| 38.145.208.206:8449 | ✓ 1357ms | ✓ 1370ms | ✓ 273ms | ✓ 1300ms | ✓ 1309ms | http |
| 38.145.220.39:8449 | ✓ 1726ms | ✓ 1498ms | ✓ 235ms | ✓ 1003ms | ✓ 1593ms | http |
| 38.145.220.182:8450 | ✓ 1244ms | 否 | ✓ 405ms | ✓ 1080ms | ✓ 1336ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1358ms | ✓ 1937ms | ✓ 1507ms | http |
| 45.186.6.104:3128 | ✓ 1421ms | ✓ 1963ms | ✓ 1957ms | 否 | 否 | http |
| 38.145.203.132:8450 | ✓ 1381ms | 否 | ✓ 1366ms | ✓ 1295ms | ✓ 1188ms | http |
| 38.145.208.210:8448 | ✓ 213ms | ✓ 827ms | ✓ 368ms | ✓ 1249ms | ✓ 1366ms | http |
| 38.145.208.235:8447 | ✓ 1134ms | 否 | ✓ 359ms | ✓ 876ms | ✓ 771ms | http |
| 218.60.0.214:80 | ✓ 985ms | ✓ 1249ms | 否 | 否 | ✓ 1007ms | http |
| 38.145.220.65:8444 | ✓ 692ms | ✓ 1104ms | ✓ 720ms | ✓ 1730ms | ✓ 980ms | http |
| 38.145.220.55:8444 | ✓ 685ms | ✓ 1897ms | ✓ 891ms | ✓ 1454ms | ✓ 978ms | http |
| 45.136.130.251:8449 | ✓ 1173ms | 否 | ✓ 125ms | ✓ 886ms | ✓ 1308ms | http |
| 45.136.130.250:8449 | ✓ 1193ms | 否 | ✓ 123ms | ✓ 904ms | ✓ 1354ms | http |
| 45.136.130.253:8449 | ✓ 1168ms | 否 | ✓ 100ms | ✓ 856ms | 否 | http |
| 38.145.208.234:8446 | ✓ 873ms | ✓ 1677ms | ✓ 836ms | ✓ 944ms | ✓ 713ms | http |
| 38.145.208.181:8445 | ✓ 1204ms | ✓ 895ms | ✓ 155ms | ✓ 864ms | ✓ 1208ms | http |
| 38.145.208.214:8446 | ✓ 887ms | ✓ 1901ms | ✓ 770ms | ✓ 1063ms | ✓ 491ms | http |
| 38.145.208.212:8446 | ✓ 886ms | 否 | ✓ 965ms | ✓ 787ms | ✓ 518ms | http |
| 38.34.179.83:8448 | ✓ 910ms | 否 | ✓ 156ms | ✓ 804ms | ✓ 1027ms | http |
| 38.145.220.41:8444 | ✓ 860ms | ✓ 873ms | ✓ 476ms | ✓ 1272ms | ✓ 1381ms | http |
| 38.145.218.206:8444 | ✓ 874ms | ✓ 879ms | ✓ 476ms | ✓ 1446ms | ✓ 1318ms | http |
| 38.145.208.244:8446 | ✓ 1579ms | 否 | ✓ 704ms | 否 | ✓ 1376ms | http |
| 38.34.179.94:8444 | ✓ 821ms | 否 | ✓ 616ms | ✓ 1044ms | ✓ 938ms | http |
| 38.34.179.101:8453 | ✓ 260ms | 否 | ✓ 184ms | ✓ 946ms | ✓ 534ms | http |
| 38.34.179.87:8447 | ✓ 246ms | ✓ 1978ms | ✓ 123ms | ✓ 778ms | ✓ 788ms | http |
| 38.34.179.86:8447 | ✓ 247ms | 否 | ✓ 158ms | ✓ 763ms | ✓ 782ms | http |
| 38.34.179.89:8444 | ✓ 496ms | ✓ 929ms | ✓ 134ms | ✓ 747ms | ✓ 989ms | http |
| 38.34.179.85:8444 | ✓ 478ms | ✓ 1890ms | ✓ 555ms | ✓ 1323ms | ✓ 638ms | http |
| 38.34.179.24:8447 | ✓ 578ms | 否 | ✓ 186ms | ✓ 703ms | ✓ 506ms | http |
| 38.145.220.79:8450 | ✓ 1265ms | 否 | ✓ 552ms | ✓ 1441ms | ✓ 1014ms | http |
| 38.145.220.43:8450 | ✓ 1207ms | 否 | ✓ 749ms | ✓ 1459ms | ✓ 922ms | http |
| 139.159.99.242:8080 | ✓ 988ms | ✓ 984ms | ✓ 842ms | ✓ 1154ms | 否 | http |
| 38.145.218.160:8448 | ✓ 1587ms | ✓ 1069ms | ✓ 309ms | ✓ 740ms | ✓ 532ms | http |
| 38.34.179.97:8446 | ✓ 1628ms | ✓ 1616ms | ✓ 289ms | ✓ 1311ms | ✓ 533ms | http |
| 45.136.131.27:8444 | ✓ 1608ms | ✓ 972ms | ✓ 836ms | ✓ 956ms | ✓ 1092ms | http |
| 38.34.179.105:8447 | ✓ 1614ms | ✓ 1061ms | ✓ 740ms | ✓ 726ms | ✓ 543ms | http |
| 3.99.169.21:13682 | ✓ 1417ms | 否 | ✓ 1708ms | 否 | ✓ 1886ms | http |
| 38.34.179.22:8449 | ✓ 1189ms | 否 | ✓ 700ms | 否 | ✓ 646ms | http |
| 45.136.131.36:8450 | ✓ 952ms | ✓ 1132ms | ✓ 1166ms | 否 | ✓ 822ms | http |
| 38.145.218.162:8448 | ✓ 1775ms | 否 | ✓ 1206ms | 否 | ✓ 1779ms | http |
| 147.45.136.99:3128 | 否 | 否 | ✓ 1427ms | ✓ 1892ms | ✓ 1816ms | http |
| 45.149.92.147:5001 | ✓ 606ms | 否 | ✓ 596ms | ✓ 955ms | ✓ 615ms | http |
| 160.20.38.102:8080 | ✓ 1803ms | 否 | 否 | ✓ 1319ms | ✓ 1290ms | http |
| 103.82.23.118:5224 | ✓ 1438ms | 否 | ✓ 1274ms | 否 | ✓ 1404ms | http |
| 103.236.142.41:8181 | ✓ 1209ms | 否 | ✓ 1342ms | ✓ 1530ms | ✓ 1328ms | http |
| 38.34.178.245:8446 | ✓ 365ms | ✓ 797ms | ✓ 999ms | ✓ 1630ms | ✓ 834ms | http |
| 38.145.218.113:8446 | ✓ 350ms | ✓ 1877ms | ✓ 1111ms | ✓ 1387ms | ✓ 585ms | http |
| 212.58.132.5:8888 | ✓ 1375ms | 否 | ✓ 1607ms | ✓ 1687ms | ✓ 1525ms | http |
| 45.140.147.155:1082 | ✓ 589ms | 否 | ✓ 609ms | ✓ 1843ms | ✓ 1164ms | http |
| 45.140.147.155:1081 | ✓ 605ms | 否 | ✓ 603ms | ✓ 1844ms | ✓ 1170ms | http |
| 61.52.131.172:8443 | ✓ 903ms | ✓ 1093ms | ✓ 865ms | ✓ 1177ms | ✓ 895ms | http |

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
