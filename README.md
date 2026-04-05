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

最后更新：2026-04-05 15:33:45 UTC（2026-04-05 23:33:45 UTC+8）

**代理总数：85**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 815ms | ✓ 1026ms | ✓ 847ms | ✓ 1060ms | ✓ 892ms | http |
| 147.161.210.140:8800 | ✓ 663ms | 否 | ✓ 985ms | ✓ 1788ms | ✓ 1117ms | http |
| 159.223.71.162:8080 | ✓ 753ms | 否 | ✓ 1594ms | ✓ 1067ms | ✓ 1347ms | http |
| 167.103.115.102:8800 | ✓ 1096ms | 否 | ✓ 1088ms | ✓ 1246ms | ✓ 1615ms | http |
| 113.160.132.26:8080 | ✓ 1924ms | ✓ 1310ms | ✓ 1300ms | ✓ 1223ms | ✓ 1222ms | http |
| 167.103.34.108:8800 | ✓ 1522ms | 否 | ✓ 1540ms | 否 | ✓ 1447ms | http |
| 111.227.254.12:22222 | ✓ 1733ms | ✓ 1402ms | ✓ 1086ms | 否 | ✓ 1804ms | http |
| 111.227.254.9:22222 | ✓ 1900ms | ✓ 1358ms | ✓ 1753ms | ✓ 1694ms | 否 | http |
| 111.227.254.11:22222 | ✓ 1056ms | ✓ 1922ms | ✓ 1782ms | 否 | 否 | http |
| 45.167.124.52:8080 | ✓ 636ms | 否 | ✓ 815ms | ✓ 1720ms | ✓ 1355ms | http |
| 45.167.125.21:999 | ✓ 1041ms | 否 | ✓ 1208ms | ✓ 1849ms | ✓ 1689ms | http |
| 167.103.144.127:8800 | ✓ 1827ms | 否 | ✓ 991ms | ✓ 1217ms | ✓ 1093ms | http |
| 167.103.31.122:8800 | ✓ 1855ms | 否 | ✓ 1429ms | ✓ 1677ms | ✓ 1588ms | http |
| 209.38.154.7:1080 | ✓ 655ms | ✓ 1737ms | ✓ 1225ms | ✓ 1782ms | 否 | http |
| 159.223.71.162:443 | ✓ 1470ms | 否 | ✓ 887ms | ✓ 1035ms | ✓ 855ms | http |
| 101.43.127.100:8877 | ✓ 863ms | ✓ 1017ms | ✓ 843ms | 否 | ✓ 926ms | http |
| 34.101.184.164:3128 | ✓ 1597ms | 否 | ✓ 1120ms | ✓ 1751ms | ✓ 1230ms | http |
| 147.161.239.240:8800 | ✓ 1381ms | 否 | 否 | ✓ 1612ms | ✓ 1561ms | http |
| 111.227.254.10:22222 | 否 | ✓ 1357ms | ✓ 1104ms | 否 | ✓ 1104ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1930ms | ✓ 1778ms | ✓ 1982ms | http |
| 64.227.76.27:1080 | ✓ 870ms | 否 | ✓ 1648ms | 否 | ✓ 1952ms | http |
| 38.145.208.185:8453 | ✓ 769ms | 否 | ✓ 401ms | ✓ 1092ms | ✓ 1199ms | http |
| 38.34.179.186:8444 | ✓ 987ms | ✓ 1858ms | 否 | ✓ 801ms | ✓ 524ms | http |
| 38.34.179.6:8449 | ✓ 778ms | ✓ 1455ms | 否 | ✓ 1473ms | 否 | http |
| 38.34.179.18:8451 | ✓ 329ms | ✓ 695ms | ✓ 1107ms | ✓ 1399ms | ✓ 547ms | http |
| 120.92.212.16:8890 | ✓ 976ms | ✓ 1207ms | ✓ 1111ms | ✓ 1243ms | ✓ 967ms | http |
| 45.136.130.178:8453 | ✓ 735ms | ✓ 1688ms | 否 | ✓ 1077ms | 否 | http |
| 120.92.212.16:7890 | ✓ 943ms | ✓ 1209ms | ✓ 1161ms | ✓ 1243ms | ✓ 1003ms | http |
| 35.225.22.61:80 | 否 | ✓ 1828ms | ✓ 1369ms | ✓ 1247ms | 否 | http |
| 20.2.83.243:3128 | ✓ 705ms | 否 | ✓ 922ms | ✓ 921ms | ✓ 708ms | http |
| 43.167.237.94:3128 | ✓ 1928ms | 否 | ✓ 1554ms | 否 | ✓ 1955ms | http |
| 43.224.171.232:8182 | ✓ 1847ms | 否 | ✓ 1562ms | ✓ 1704ms | ✓ 1614ms | http |
| 148.251.87.79:16379 | ✓ 1268ms | 否 | ✓ 1795ms | 否 | ✓ 1762ms | http |
| 45.136.131.25:8453 | ✓ 927ms | ✓ 1281ms | ✓ 1231ms | ✓ 989ms | ✓ 971ms | http |
| 210.223.44.230:3128 | ✓ 1399ms | ✓ 873ms | ✓ 1008ms | ✓ 1124ms | ✓ 1341ms | http |
| 38.34.179.150:8449 | ✓ 585ms | ✓ 721ms | ✓ 275ms | ✓ 664ms | ✓ 593ms | http |
| 45.136.130.193:8444 | ✓ 688ms | 否 | ✓ 170ms | ✓ 872ms | ✓ 1006ms | http |
| 1.231.81.166:3128 | ✓ 784ms | ✓ 1421ms | 否 | ✓ 1063ms | ✓ 950ms | http |
| 38.34.179.184:8448 | ✓ 167ms | ✓ 772ms | ✓ 200ms | ✓ 893ms | ✓ 671ms | http |
| 38.34.179.47:8452 | ✓ 954ms | 否 | ✓ 131ms | ✓ 691ms | ✓ 533ms | http |
| 45.136.131.63:8443 | ✓ 661ms | ✓ 691ms | ✓ 89ms | ✓ 859ms | ✓ 741ms | http |
| 38.34.179.213:8452 | ✓ 870ms | ✓ 1136ms | ✓ 101ms | ✓ 862ms | ✓ 1015ms | http |
| 47.74.226.8:5001 | 否 | 否 | ✓ 1056ms | ✓ 1287ms | ✓ 1489ms | http |
| 38.34.179.89:8451 | 否 | ✓ 1297ms | ✓ 80ms | ✓ 696ms | ✓ 638ms | http |
| 38.34.179.87:8451 | 否 | 否 | ✓ 170ms | ✓ 934ms | ✓ 1482ms | http |
| 45.136.131.28:8447 | ✓ 1693ms | 否 | ✓ 194ms | 否 | ✓ 1300ms | http |
| 38.145.208.211:8453 | ✓ 1989ms | 否 | ✓ 445ms | ✓ 799ms | ✓ 530ms | http |
| 45.136.130.246:8445 | 否 | 否 | ✓ 1584ms | ✓ 1660ms | ✓ 987ms | http |
| 38.145.203.45:8451 | ✓ 508ms | ✓ 1593ms | ✓ 737ms | ✓ 1101ms | ✓ 1830ms | http |
| 38.145.218.51:8444 | ✓ 983ms | 否 | ✓ 105ms | ✓ 1144ms | ✓ 1287ms | http |
| 38.145.218.14:8446 | ✓ 1863ms | ✓ 1820ms | 否 | ✓ 1867ms | 否 | http |
| 38.145.203.35:8450 | ✓ 1276ms | 否 | ✓ 467ms | ✓ 1662ms | 否 | http |
| 38.145.218.9:8445 | ✓ 100ms | ✓ 1021ms | ✓ 131ms | ✓ 793ms | ✓ 800ms | http |
| 38.34.179.155:8453 | ✓ 1418ms | ✓ 1114ms | ✓ 92ms | ✓ 668ms | ✓ 534ms | http |
| 38.34.179.81:8443 | ✓ 208ms | ✓ 679ms | ✓ 299ms | ✓ 1642ms | ✓ 1890ms | http |
| 38.145.208.243:8447 | ✓ 220ms | ✓ 802ms | ✓ 202ms | ✓ 1228ms | ✓ 1661ms | http |
| 5.104.87.17:8051 | ✓ 1531ms | 否 | ✓ 1475ms | ✓ 1646ms | 否 | http |
| 185.114.73.2:1080 | 否 | 否 | ✓ 1868ms | ✓ 1795ms | ✓ 1617ms | http |
| 208.87.243.199:7878 | ✓ 1398ms | ✓ 770ms | ✓ 795ms | ✓ 1081ms | ✓ 682ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1254ms | ✓ 1585ms | ✓ 1060ms | http |
| 194.67.99.223:1080 | ✓ 1146ms | 否 | 否 | ✓ 1985ms | ✓ 1487ms | http |
| 38.34.183.222:8453 | 否 | ✓ 814ms | ✓ 219ms | ✓ 702ms | ✓ 577ms | http |
| 103.113.70.189:1081 | ✓ 432ms | ✓ 1779ms | 否 | ✓ 1363ms | ✓ 1517ms | http |
| 150.107.140.238:3128 | ✓ 1574ms | 否 | ✓ 1032ms | 否 | ✓ 903ms | http |
| 86.53.183.16:1080 | ✓ 1049ms | 否 | ✓ 928ms | 否 | ✓ 1620ms | http |
| 202.154.18.80:8082 | ✓ 1075ms | 否 | 否 | ✓ 1490ms | ✓ 1516ms | http |
| 38.145.220.60:8447 | 否 | 否 | ✓ 1528ms | ✓ 1368ms | ✓ 1907ms | http |
| 45.136.131.35:8452 | 否 | ✓ 847ms | ✓ 886ms | ✓ 717ms | ✓ 570ms | http |
| 38.145.208.220:8448 | 否 | ✓ 1229ms | ✓ 725ms | ✓ 1084ms | ✓ 744ms | http |
| 45.136.131.54:8448 | 否 | ✓ 1615ms | ✓ 446ms | ✓ 865ms | ✓ 1334ms | http |
| 38.145.208.217:8450 | 否 | ✓ 1320ms | ✓ 608ms | ✓ 1430ms | ✓ 705ms | http |
| 45.136.130.188:8449 | 否 | ✓ 981ms | ✓ 754ms | 否 | ✓ 562ms | http |
| 45.136.131.46:8443 | 否 | ✓ 1493ms | ✓ 630ms | ✓ 934ms | ✓ 1075ms | http |
| 38.34.179.162:8444 | 否 | ✓ 1949ms | ✓ 834ms | ✓ 1001ms | ✓ 809ms | http |
| 38.34.179.156:8444 | 否 | 否 | ✓ 820ms | ✓ 987ms | ✓ 805ms | http |
| 38.145.220.198:8446 | 否 | ✓ 1383ms | ✓ 1280ms | ✓ 935ms | 否 | http |
| 38.34.183.13:8449 | 否 | 否 | ✓ 1300ms | ✓ 1347ms | ✓ 1480ms | http |
| 103.179.252.81:3128 | ✓ 1828ms | 否 | ✓ 1317ms | ✓ 1664ms | ✓ 1446ms | http |
| 150.241.71.15:1080 | ✓ 747ms | 否 | ✓ 1271ms | ✓ 1623ms | 否 | http |
| 183.232.248.73:7890 | ✓ 890ms | ✓ 1109ms | ✓ 1144ms | ✓ 1039ms | ✓ 855ms | http |
| 133.242.138.34:8100 | ✓ 1491ms | 否 | ✓ 1113ms | ✓ 899ms | 否 | http |
| 101.132.61.121:8888 | ✓ 1270ms | ✓ 1203ms | ✓ 1279ms | ✓ 1373ms | ✓ 1247ms | http |
| 61.52.131.172:8443 | ✓ 968ms | ✓ 1101ms | ✓ 896ms | ✓ 1119ms | ✓ 914ms | http |
| 45.136.131.39:8451 | ✓ 1281ms | 否 | ✓ 1042ms | ✓ 678ms | ✓ 524ms | http |
| 103.101.193.170:1111 | 否 | 否 | ✓ 1466ms | ✓ 1360ms | ✓ 1485ms | http |

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
