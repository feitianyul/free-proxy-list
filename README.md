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

最后更新：2026-04-02 18:53:02 UTC（2026-04-03 02:53:02 UTC+8）

**代理总数：152**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 152 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 152 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 220ms | ✓ 659ms | ✓ 990ms | ✓ 750ms | ✓ 657ms | http |
| 35.225.22.61:80 | ✓ 600ms | ✓ 1524ms | 否 | ✓ 1108ms | ✓ 942ms | http |
| 203.80.138.81:50000 | ✓ 900ms | ✓ 1126ms | ✓ 990ms | ✓ 986ms | ✓ 808ms | http |
| 43.99.54.236:5555 | ✓ 669ms | ✓ 1612ms | ✓ 671ms | ✓ 782ms | ✓ 633ms | http |
| 147.161.210.140:8800 | ✓ 1475ms | ✓ 1059ms | ✓ 813ms | ✓ 1052ms | ✓ 980ms | http |
| 1.231.81.166:3128 | ✓ 1504ms | ✓ 921ms | ✓ 1679ms | ✓ 1212ms | ✓ 813ms | http |
| 159.223.71.162:8080 | ✓ 700ms | 否 | 否 | ✓ 1022ms | ✓ 912ms | http |
| 167.103.115.102:8800 | ✓ 947ms | 否 | ✓ 1319ms | ✓ 1256ms | ✓ 1144ms | http |
| 113.160.132.26:8080 | ✓ 1482ms | ✓ 1402ms | ✓ 978ms | ✓ 1662ms | ✓ 928ms | http |
| 147.161.239.240:8800 | ✓ 995ms | ✓ 1814ms | ✓ 1125ms | ✓ 1873ms | ✓ 1587ms | http |
| 167.103.34.108:8800 | ✓ 1344ms | 否 | ✓ 1498ms | ✓ 1796ms | 否 | http |
| 95.213.217.168:52004 | ✓ 1608ms | ✓ 1852ms | 否 | 否 | ✓ 1709ms | http |
| 212.58.132.5:8888 | ✓ 1137ms | 否 | ✓ 1591ms | 否 | ✓ 1227ms | http |
| 45.167.124.52:8080 | ✓ 1241ms | 否 | ✓ 1449ms | 否 | ✓ 1981ms | http |
| 38.145.220.102:8453 | ✓ 1113ms | ✓ 898ms | ✓ 940ms | ✓ 1091ms | ✓ 1294ms | http |
| 38.145.220.60:8447 | ✓ 1141ms | ✓ 1497ms | ✓ 269ms | ✓ 1192ms | 否 | http |
| 34.96.238.40:8080 | ✓ 801ms | ✓ 1228ms | 否 | ✓ 945ms | 否 | http |
| 222.184.48.242:22222 | 否 | 否 | ✓ 805ms | ✓ 1070ms | ✓ 870ms | http |
| 167.103.144.127:8800 | ✓ 1637ms | 否 | ✓ 1311ms | ✓ 1578ms | 否 | http |
| 45.140.147.82:1081 | ✓ 616ms | 否 | ✓ 1578ms | ✓ 1809ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1515ms | 否 | ✓ 1702ms | 否 | ✓ 1855ms | http |
| 45.167.125.21:999 | ✓ 1212ms | 否 | ✓ 1471ms | 否 | ✓ 1830ms | http |
| 115.231.181.40:8128 | ✓ 899ms | ✓ 1050ms | ✓ 881ms | ✓ 1215ms | ✓ 950ms | http |
| 165.232.146.249:3128 | ✓ 989ms | 否 | ✓ 1929ms | 否 | ✓ 1124ms | http |
| 164.90.151.28:3128 | 否 | 否 | ✓ 1948ms | ✓ 1279ms | ✓ 1548ms | http |
| 180.250.219.58:53281 | ✓ 1780ms | 否 | ✓ 1470ms | ✓ 1907ms | 否 | http |
| 222.184.48.251:22222 | ✓ 1180ms | 否 | ✓ 1055ms | ✓ 1590ms | ✓ 1423ms | http |
| 45.12.151.226:2829 | 否 | ✓ 1965ms | ✓ 1422ms | 否 | ✓ 1620ms | http |
| 38.34.179.98:8451 | ✓ 742ms | ✓ 859ms | ✓ 833ms | ✓ 808ms | ✓ 516ms | http |
| 34.101.184.164:3128 | ✓ 1930ms | 否 | ✓ 1074ms | 否 | ✓ 1385ms | http |
| 59.46.216.131:30001 | ✓ 1987ms | ✓ 1300ms | ✓ 1998ms | 否 | ✓ 1021ms | http |
| 38.34.183.225:8450 | ✓ 276ms | ✓ 761ms | ✓ 231ms | 否 | 否 | http |
| 38.145.203.19:8447 | ✓ 427ms | ✓ 880ms | ✓ 214ms | ✓ 956ms | ✓ 587ms | http |
| 38.145.220.11:8447 | ✓ 412ms | ✓ 658ms | ✓ 614ms | ✓ 721ms | ✓ 676ms | http |
| 45.136.130.177:8448 | ✓ 411ms | ✓ 665ms | ✓ 175ms | ✓ 775ms | ✓ 834ms | http |
| 128.199.116.219:9090 | ✓ 767ms | 否 | ✓ 976ms | ✓ 1094ms | ✓ 822ms | http |
| 128.199.121.61:9090 | ✓ 698ms | 否 | ✓ 759ms | ✓ 1022ms | ✓ 820ms | http |
| 101.43.127.100:8877 | ✓ 1527ms | ✓ 1033ms | ✓ 839ms | ✓ 1043ms | ✓ 829ms | http |
| 159.223.71.162:443 | ✓ 902ms | 否 | ✓ 887ms | ✓ 1044ms | ✓ 884ms | http |
| 120.92.212.16:7890 | ✓ 922ms | 否 | 否 | ✓ 1361ms | ✓ 945ms | http |
| 160.250.5.22:1 | 否 | 否 | ✓ 1301ms | ✓ 1562ms | ✓ 1116ms | http |
| 38.145.203.97:8448 | ✓ 411ms | ✓ 1123ms | ✓ 1609ms | ✓ 1082ms | ✓ 647ms | http |
| 31.192.106.135:8005 | ✓ 1604ms | 否 | ✓ 1720ms | 否 | ✓ 1965ms | http |
| 38.145.203.106:8448 | ✓ 411ms | ✓ 1127ms | ✓ 1610ms | ✓ 1291ms | ✓ 785ms | http |
| 177.234.217.88:999 | ✓ 1736ms | 否 | ✓ 1852ms | 否 | ✓ 1916ms | http |
| 38.34.179.228:8453 | ✓ 680ms | ✓ 765ms | ✓ 379ms | ✓ 1042ms | ✓ 1745ms | http |
| 38.34.179.61:8445 | ✓ 749ms | ✓ 1097ms | ✓ 707ms | ✓ 811ms | ✓ 745ms | http |
| 8.219.97.248:80 | ✓ 1573ms | 否 | ✓ 1565ms | ✓ 1636ms | 否 | http |
| 38.34.183.130:8452 | ✓ 872ms | ✓ 1902ms | ✓ 896ms | ✓ 1586ms | 否 | http |
| 38.34.183.224:8448 | ✓ 228ms | ✓ 791ms | ✓ 257ms | ✓ 1125ms | 否 | http |
| 178.128.243.121:3128 | ✓ 793ms | 否 | ✓ 1891ms | ✓ 1967ms | ✓ 1525ms | http |
| 150.249.255.91:3128 | 否 | ✓ 1101ms | ✓ 941ms | ✓ 798ms | ✓ 651ms | http |
| 38.34.179.14:8450 | ✓ 779ms | ✓ 647ms | ✓ 240ms | ✓ 980ms | ✓ 1903ms | http |
| 118.31.1.154:80 | ✓ 829ms | ✓ 1398ms | ✓ 853ms | ✓ 1116ms | ✓ 962ms | http |
| 157.0.142.246:10061 | ✓ 1197ms | ✓ 1268ms | ✓ 998ms | ✓ 1322ms | ✓ 1074ms | http |
| 133.242.138.34:8100 | ✓ 1526ms | ✓ 1611ms | ✓ 1007ms | ✓ 1416ms | ✓ 1301ms | http |
| 195.123.209.48:3128 | ✓ 1209ms | 否 | ✓ 1752ms | 否 | ✓ 1859ms | http |
| 222.184.48.252:22222 | ✓ 1291ms | ✓ 1722ms | ✓ 1659ms | ✓ 1928ms | 否 | http |
| 45.140.147.82:1082 | ✓ 1193ms | ✓ 1466ms | 否 | 否 | ✓ 1021ms | http |
| 116.80.49.163:3172 | 否 | 否 | ✓ 1460ms | ✓ 1870ms | ✓ 1774ms | http |
| 38.145.208.242:8451 | ✓ 509ms | ✓ 994ms | ✓ 881ms | 否 | ✓ 495ms | http |
| 45.140.147.155:1081 | ✓ 622ms | ✓ 1499ms | ✓ 998ms | ✓ 1611ms | 否 | http |
| 38.34.183.164:8444 | 否 | ✓ 658ms | ✓ 323ms | ✓ 1083ms | ✓ 1533ms | http |
| 42.96.16.158:1311 | ✓ 1540ms | 否 | ✓ 1636ms | ✓ 1190ms | ✓ 1374ms | http |
| 45.140.147.155:1082 | ✓ 722ms | ✓ 1489ms | 否 | 否 | ✓ 1240ms | http |
| 38.34.179.27:8451 | 否 | 否 | ✓ 78ms | ✓ 694ms | ✓ 813ms | http |
| 120.92.212.16:8890 | ✓ 1097ms | 否 | ✓ 1778ms | 否 | ✓ 1285ms | http |
| 5.102.109.41:999 | ✓ 422ms | 否 | ✓ 558ms | ✓ 1268ms | ✓ 1254ms | http |
| 139.135.79.91:8082 | ✓ 1464ms | 否 | ✓ 1686ms | ✓ 1588ms | 否 | http |
| 121.230.8.45:1080 | ✓ 1070ms | ✓ 1484ms | ✓ 1020ms | ✓ 1574ms | ✓ 1226ms | http |
| 121.230.8.22:1080 | ✓ 1215ms | ✓ 1586ms | ✓ 1101ms | ✓ 1344ms | ✓ 1255ms | http |
| 38.34.179.186:8444 | ✓ 1913ms | ✓ 833ms | ✓ 488ms | ✓ 1202ms | 否 | http |
| 128.199.114.189:9090 | ✓ 1014ms | 否 | ✓ 1137ms | ✓ 1008ms | ✓ 935ms | http |
| 128.199.113.85:9090 | ✓ 1031ms | 否 | ✓ 1154ms | ✓ 1027ms | ✓ 869ms | http |
| 38.145.220.198:8448 | ✓ 275ms | ✓ 675ms | ✓ 1134ms | ✓ 1242ms | ✓ 708ms | http |
| 91.233.223.147:3128 | ✓ 1300ms | 否 | ✓ 1516ms | 否 | ✓ 1820ms | http |
| 45.136.131.61:8447 | ✓ 1558ms | 否 | ✓ 763ms | ✓ 1011ms | 否 | http |
| 38.145.218.101:8447 | 否 | 否 | ✓ 698ms | ✓ 895ms | ✓ 941ms | http |
| 38.145.208.179:8447 | ✓ 204ms | ✓ 665ms | ✓ 116ms | ✓ 812ms | ✓ 725ms | http |
| 38.34.179.35:8448 | ✓ 173ms | ✓ 860ms | ✓ 1669ms | ✓ 1267ms | ✓ 725ms | http |
| 38.145.208.182:8450 | ✓ 672ms | ✓ 677ms | ✓ 149ms | ✓ 648ms | ✓ 501ms | http |
| 38.145.203.46:8448 | ✓ 695ms | ✓ 673ms | ✓ 133ms | ✓ 694ms | ✓ 512ms | http |
| 38.145.203.45:8452 | ✓ 693ms | ✓ 673ms | ✓ 136ms | ✓ 684ms | ✓ 518ms | http |
| 45.136.130.166:8449 | ✓ 673ms | ✓ 943ms | ✓ 75ms | ✓ 690ms | ✓ 681ms | http |
| 38.34.179.185:8445 | ✓ 693ms | ✓ 710ms | ✓ 164ms | ✓ 883ms | ✓ 655ms | http |
| 38.145.208.227:8451 | ✓ 694ms | ✓ 941ms | ✓ 1296ms | ✓ 762ms | ✓ 539ms | http |
| 45.136.130.180:8453 | ✓ 673ms | ✓ 741ms | ✓ 335ms | ✓ 1142ms | ✓ 1509ms | http |
| 38.145.220.188:8444 | ✓ 672ms | ✓ 795ms | ✓ 211ms | ✓ 955ms | ✓ 1197ms | http |
| 38.34.179.202:8449 | ✓ 701ms | ✓ 620ms | ✓ 187ms | ✓ 957ms | 否 | http |
| 45.136.130.171:8445 | ✓ 673ms | ✓ 1067ms | ✓ 416ms | ✓ 1246ms | ✓ 527ms | http |
| 38.34.179.68:8452 | ✓ 992ms | ✓ 943ms | ✓ 508ms | ✓ 978ms | ✓ 706ms | http |
| 38.34.179.100:8449 | ✓ 695ms | ✓ 843ms | ✓ 119ms | ✓ 665ms | ✓ 738ms | http |
| 45.136.130.248:8452 | ✓ 1513ms | ✓ 657ms | ✓ 126ms | ✓ 982ms | 否 | http |
| 38.34.179.21:8446 | ✓ 992ms | ✓ 616ms | ✓ 1455ms | ✓ 1561ms | ✓ 865ms | http |
| 45.136.130.169:8446 | ✓ 676ms | ✓ 1036ms | ✓ 321ms | ✓ 1015ms | ✓ 532ms | http |
| 38.145.208.241:8453 | ✓ 673ms | ✓ 1494ms | ✓ 809ms | ✓ 655ms | ✓ 868ms | http |
| 38.145.203.135:8444 | ✓ 672ms | ✓ 1342ms | ✓ 588ms | ✓ 673ms | ✓ 676ms | http |
| 38.34.179.179:8449 | ✓ 693ms | ✓ 855ms | ✓ 1019ms | ✓ 701ms | ✓ 573ms | http |
| 38.34.179.77:8445 | ✓ 675ms | 否 | ✓ 84ms | ✓ 663ms | ✓ 637ms | http |
| 38.34.179.16:8451 | ✓ 695ms | ✓ 798ms | ✓ 1253ms | ✓ 1082ms | ✓ 877ms | http |

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
