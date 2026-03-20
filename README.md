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

最后更新：2026-03-20 19:46:34 UTC（2026-03-21 03:46:34 UTC+8）

**代理总数：90**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 90 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | ✓ 1170ms | ✓ 1893ms | ✓ 1515ms | ✓ 1167ms | ✓ 1399ms | http |
| 147.161.210.140:8800 | ✓ 1440ms | 否 | ✓ 762ms | ✓ 1306ms | ✓ 1144ms | http |
| 178.156.187.185:10001 | ✓ 505ms | ✓ 1367ms | ✓ 1934ms | ✓ 1975ms | ✓ 1574ms | http |
| 147.161.239.240:8800 | ✓ 1242ms | ✓ 1748ms | ✓ 1138ms | ✓ 1802ms | ✓ 1532ms | http |
| 174.138.24.77:1080 | ✓ 958ms | 否 | ✓ 1594ms | ✓ 1131ms | ✓ 1111ms | http |
| 113.160.132.26:8080 | ✓ 1682ms | 否 | ✓ 1260ms | 否 | ✓ 1372ms | http |
| 45.167.124.52:8080 | ✓ 1695ms | ✓ 1951ms | ✓ 1243ms | ✓ 1743ms | ✓ 1594ms | http |
| 43.99.54.236:5555 | ✓ 636ms | ✓ 1127ms | ✓ 631ms | ✓ 810ms | ✓ 639ms | http |
| 219.117.204.211:7799 | ✓ 478ms | ✓ 1121ms | ✓ 483ms | ✓ 835ms | 否 | http |
| 137.220.150.22:6005 | ✓ 1397ms | 否 | ✓ 963ms | ✓ 1150ms | ✓ 881ms | http |
| 167.103.34.108:8800 | ✓ 1419ms | 否 | ✓ 1163ms | ✓ 1335ms | ✓ 1270ms | http |
| 8.219.97.248:80 | ✓ 1599ms | 否 | ✓ 1298ms | ✓ 1745ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1707ms | 否 | ✓ 1830ms | ✓ 1948ms | ✓ 1862ms | http |
| 137.220.150.152:6005 | ✓ 1347ms | 否 | ✓ 890ms | ✓ 1264ms | 否 | http |
| 186.148.180.46:999 | 否 | 否 | ✓ 1276ms | ✓ 1659ms | ✓ 1518ms | http |
| 115.231.181.40:8128 | ✓ 1207ms | 否 | ✓ 848ms | ✓ 1170ms | ✓ 905ms | http |
| 38.145.218.212:8448 | ✓ 497ms | ✓ 646ms | ✓ 270ms | ✓ 697ms | ✓ 1022ms | http |
| 101.43.127.100:8877 | ✓ 945ms | ✓ 1023ms | ✓ 901ms | ✓ 1062ms | ✓ 865ms | http |
| 120.92.212.16:7890 | ✓ 953ms | 否 | ✓ 1148ms | 否 | ✓ 952ms | http |
| 91.238.105.64:2024 | ✓ 1440ms | 否 | ✓ 1913ms | 否 | ✓ 1920ms | http |
| 162.240.154.26:3128 | ✓ 1660ms | 否 | ✓ 1027ms | 否 | ✓ 1160ms | http |
| 45.136.130.252:8451 | ✓ 411ms | ✓ 801ms | ✓ 586ms | ✓ 753ms | ✓ 684ms | http |
| 38.34.179.58:8447 | ✓ 610ms | 否 | 否 | ✓ 998ms | ✓ 1060ms | http |
| 194.67.99.223:1080 | ✓ 1175ms | 否 | ✓ 1515ms | 否 | ✓ 1860ms | http |
| 45.136.130.186:8451 | ✓ 839ms | ✓ 593ms | ✓ 431ms | ✓ 673ms | ✓ 989ms | http |
| 1.231.81.166:3128 | ✓ 1619ms | ✓ 1244ms | 否 | ✓ 1048ms | ✓ 1223ms | http |
| 34.96.238.40:8080 | ✓ 1219ms | ✓ 1207ms | 否 | 否 | ✓ 1295ms | http |
| 137.220.150.104:6005 | ✓ 1140ms | 否 | ✓ 910ms | ✓ 1436ms | 否 | http |
| 120.92.212.16:8890 | ✓ 945ms | ✓ 1192ms | 否 | ✓ 1262ms | 否 | http |
| 137.220.151.110:6005 | 否 | 否 | ✓ 1610ms | ✓ 1184ms | ✓ 1172ms | http |
| 38.34.179.16:8451 | ✓ 1088ms | ✓ 1832ms | ✓ 211ms | ✓ 858ms | ✓ 1513ms | http |
| 222.228.171.92:8080 | ✓ 1584ms | 否 | ✓ 1376ms | ✓ 803ms | ✓ 735ms | http |
| 59.46.216.131:30001 | ✓ 1705ms | ✓ 1297ms | ✓ 1104ms | 否 | 否 | http |
| 222.127.55.155:8082 | ✓ 1620ms | 否 | 否 | ✓ 1308ms | ✓ 1578ms | http |
| 182.253.93.3:53281 | ✓ 1795ms | 否 | 否 | ✓ 1545ms | ✓ 1426ms | http |
| 103.111.136.82:34564 | 否 | 否 | ✓ 1661ms | ✓ 1575ms | ✓ 1568ms | http |
| 38.34.179.61:8445 | ✓ 300ms | ✓ 756ms | ✓ 891ms | ✓ 1665ms | ✓ 663ms | http |
| 210.223.44.230:3128 | ✓ 577ms | ✓ 1132ms | 否 | 否 | ✓ 668ms | http |
| 38.34.183.233:8448 | 否 | ✓ 712ms | ✓ 722ms | ✓ 1014ms | ✓ 1420ms | http |
| 137.220.150.170:6005 | ✓ 1045ms | ✓ 1876ms | ✓ 797ms | ✓ 1722ms | ✓ 926ms | http |
| 38.34.179.14:8450 | 否 | ✓ 887ms | ✓ 1119ms | ✓ 1755ms | ✓ 614ms | http |
| 5.102.109.41:999 | ✓ 615ms | ✓ 1271ms | 否 | 否 | ✓ 1309ms | http |
| 103.46.11.156:3125 | ✓ 1938ms | 否 | ✓ 1628ms | 否 | ✓ 1458ms | http |
| 45.136.131.35:8452 | ✓ 300ms | ✓ 694ms | ✓ 614ms | ✓ 729ms | ✓ 529ms | http |
| 14.225.212.37:7890 | ✓ 1353ms | 否 | ✓ 824ms | ✓ 1768ms | ✓ 1065ms | http |
| 38.34.179.162:8451 | ✓ 873ms | ✓ 815ms | ✓ 1462ms | ✓ 1902ms | ✓ 1025ms | http |
| 38.34.183.234:8450 | ✓ 863ms | 否 | ✓ 95ms | ✓ 821ms | ✓ 1057ms | http |
| 38.145.220.198:8448 | ✓ 1884ms | ✓ 1188ms | ✓ 112ms | ✓ 840ms | ✓ 509ms | http |
| 217.76.245.80:999 | ✓ 778ms | ✓ 1561ms | ✓ 1265ms | ✓ 1533ms | ✓ 1501ms | http |
| 116.80.62.22:3128 | 否 | 否 | ✓ 1530ms | ✓ 1799ms | ✓ 1674ms | http |
| 101.32.244.83:8080 | ✓ 989ms | 否 | ✓ 947ms | ✓ 1337ms | ✓ 1218ms | http |
| 121.43.196.213:8222 | ✓ 969ms | ✓ 1059ms | ✓ 844ms | ✓ 1116ms | ✓ 874ms | http |
| 121.43.196.210:8222 | ✓ 913ms | ✓ 1090ms | ✓ 845ms | ✓ 1105ms | ✓ 922ms | http |
| 114.55.226.123:10086 | ✓ 1036ms | ✓ 1363ms | ✓ 1058ms | ✓ 1262ms | ✓ 1061ms | http |
| 35.183.64.191:44096 | ✓ 1861ms | 否 | ✓ 1985ms | 否 | ✓ 1940ms | http |
| 172.212.68.37:3128 | ✓ 577ms | ✓ 1446ms | ✓ 1679ms | ✓ 1866ms | ✓ 1260ms | http |
| 201.144.25.226:3128 | ✓ 632ms | ✓ 1200ms | ✓ 992ms | ✓ 1265ms | ✓ 941ms | http |
| 146.56.182.165:3128 | ✓ 832ms | 否 | 否 | ✓ 1488ms | ✓ 712ms | http |
| 139.159.99.242:8080 | 否 | 否 | ✓ 822ms | ✓ 1145ms | ✓ 1251ms | http |
| 157.230.220.25:4857 | ✓ 875ms | 否 | ✓ 978ms | ✓ 1383ms | 否 | http |
| 45.149.92.147:5001 | ✓ 1314ms | 否 | ✓ 626ms | ✓ 855ms | ✓ 631ms | http |
| 38.145.220.33:8448 | ✓ 663ms | 否 | ✓ 401ms | ✓ 1265ms | 否 | http |
| 103.84.95.54:7890 | ✓ 809ms | 否 | ✓ 661ms | ✓ 830ms | ✓ 706ms | http |
| 45.136.130.177:8448 | ✓ 1021ms | ✓ 1775ms | ✓ 322ms | ✓ 906ms | ✓ 1982ms | http |
| 142.171.224.229:7890 | ✓ 456ms | ✓ 1318ms | ✓ 325ms | ✓ 665ms | ✓ 473ms | http |
| 106.117.208.101:7890 | ✓ 879ms | ✓ 1274ms | ✓ 1082ms | ✓ 1204ms | ✓ 1012ms | http |
| 38.34.179.49:8450 | ✓ 1832ms | ✓ 776ms | ✓ 1191ms | 否 | 否 | http |
| 133.242.138.34:8100 | ✓ 1093ms | ✓ 1671ms | 否 | ✓ 1671ms | 否 | http |
| 116.80.65.79:3172 | 否 | 否 | ✓ 1526ms | ✓ 1838ms | ✓ 1749ms | http |
| 103.183.10.169:3125 | ✓ 1748ms | 否 | ✓ 1619ms | ✓ 1485ms | ✓ 1404ms | http |
| 45.186.6.104:3128 | ✓ 1765ms | ✓ 1659ms | ✓ 1821ms | 否 | 否 | http |
| 103.39.51.190:8080 | ✓ 1926ms | 否 | 否 | ✓ 1968ms | ✓ 1458ms | http |
| 38.34.179.96:8451 | ✓ 150ms | ✓ 1156ms | ✓ 1798ms | ✓ 750ms | ✓ 683ms | http |
| 38.188.247.12:999 | 否 | 否 | ✓ 686ms | ✓ 1699ms | ✓ 1476ms | http |
| 38.145.218.227:8445 | 否 | 否 | ✓ 1085ms | ✓ 694ms | ✓ 525ms | http |
| 38.145.208.151:8453 | ✓ 348ms | ✓ 1147ms | ✓ 1240ms | ✓ 936ms | ✓ 559ms | http |
| 117.86.6.244:1080 | ✓ 941ms | ✓ 1209ms | ✓ 970ms | ✓ 1273ms | ✓ 971ms | http |
| 45.136.130.197:8452 | 否 | 否 | ✓ 830ms | ✓ 861ms | ✓ 895ms | http |
| 38.145.220.43:8445 | 否 | ✓ 1273ms | ✓ 1332ms | 否 | ✓ 513ms | http |
| 114.237.77.231:1080 | ✓ 985ms | ✓ 1091ms | ✓ 995ms | 否 | 否 | http |
| 45.136.131.66:8445 | ✓ 1013ms | ✓ 1592ms | 否 | ✓ 980ms | ✓ 1242ms | http |
| 222.109.119.178:3128 | 否 | 否 | ✓ 768ms | ✓ 912ms | ✓ 1644ms | http |
| 38.34.179.37:8444 | ✓ 580ms | ✓ 1459ms | ✓ 185ms | ✓ 701ms | ✓ 539ms | http |
| 38.34.179.99:8448 | ✓ 1200ms | ✓ 876ms | ✓ 839ms | 否 | ✓ 839ms | http |
| 38.34.179.97:8448 | ✓ 614ms | ✓ 1102ms | ✓ 1659ms | ✓ 1169ms | ✓ 783ms | http |
| 38.34.179.20:8445 | 否 | ✓ 612ms | ✓ 314ms | ✓ 696ms | ✓ 963ms | http |
| 45.136.130.173:8448 | 否 | 否 | ✓ 1990ms | ✓ 1443ms | ✓ 674ms | http |
| 38.145.208.172:8448 | 否 | ✓ 1104ms | ✓ 406ms | ✓ 1207ms | ✓ 1958ms | http |
| 38.34.179.27:8451 | ✓ 1850ms | ✓ 1643ms | ✓ 1141ms | ✓ 1349ms | ✓ 647ms | http |
| 38.34.179.7:8443 | ✓ 1908ms | 否 | ✓ 1742ms | ✓ 1112ms | ✓ 745ms | http |

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
