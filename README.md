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

最后更新：2026-03-21 12:18:17 UTC（2026-03-21 20:18:17 UTC+8）

**代理总数：148**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 147 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 148 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1930ms | 否 | ✓ 1018ms | ✓ 1157ms | ✓ 982ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1589ms | ✓ 1881ms | ✓ 1738ms | http |
| 38.145.220.11:8445 | ✓ 858ms | ✓ 802ms | ✓ 397ms | ✓ 818ms | ✓ 651ms | http |
| 38.34.179.40:8446 | ✓ 839ms | ✓ 800ms | ✓ 397ms | ✓ 866ms | ✓ 633ms | http |
| 38.145.208.181:8445 | ✓ 821ms | ✓ 821ms | ✓ 375ms | ✓ 824ms | ✓ 731ms | http |
| 38.34.178.154:8445 | ✓ 884ms | ✓ 782ms | ✓ 409ms | ✓ 847ms | ✓ 638ms | http |
| 38.34.179.86:8452 | ✓ 867ms | ✓ 830ms | ✓ 351ms | ✓ 885ms | ✓ 676ms | http |
| 38.34.179.88:8446 | ✓ 808ms | ✓ 753ms | ✓ 427ms | ✓ 964ms | ✓ 630ms | http |
| 38.145.208.242:8451 | ✓ 793ms | ✓ 951ms | ✓ 246ms | ✓ 1025ms | ✓ 629ms | http |
| 38.145.208.244:8448 | ✓ 786ms | ✓ 949ms | ✓ 247ms | ✓ 1065ms | ✓ 701ms | http |
| 38.34.179.165:8446 | ✓ 832ms | ✓ 1321ms | ✓ 223ms | ✓ 850ms | ✓ 618ms | http |
| 38.145.208.185:8449 | ✓ 783ms | ✓ 1265ms | ✓ 215ms | ✓ 825ms | ✓ 781ms | http |
| 38.34.179.39:8452 | ✓ 832ms | ✓ 1416ms | ✓ 215ms | ✓ 829ms | ✓ 642ms | http |
| 38.34.183.47:8452 | ✓ 866ms | ✓ 802ms | ✓ 641ms | ✓ 894ms | ✓ 669ms | http |
| 38.34.179.172:8451 | ✓ 881ms | ✓ 1404ms | ✓ 263ms | ✓ 843ms | ✓ 662ms | http |
| 45.136.130.186:8451 | ✓ 896ms | ✓ 1361ms | ✓ 215ms | ✓ 915ms | ✓ 761ms | http |
| 38.145.218.229:8450 | ✓ 1105ms | ✓ 913ms | ✓ 318ms | ✓ 1183ms | ✓ 676ms | http |
| 45.136.131.62:8449 | ✓ 785ms | ✓ 1553ms | ✓ 287ms | ✓ 872ms | ✓ 610ms | http |
| 38.34.179.173:8452 | ✓ 827ms | ✓ 1564ms | ✓ 269ms | ✓ 824ms | ✓ 618ms | http |
| 38.34.183.234:8450 | ✓ 811ms | ✓ 1765ms | ✓ 208ms | ✓ 855ms | ✓ 634ms | http |
| 45.136.131.54:8448 | ✓ 852ms | ✓ 1750ms | ✓ 200ms | ✓ 822ms | ✓ 665ms | http |
| 38.34.178.155:8448 | ✓ 785ms | ✓ 1815ms | ✓ 202ms | ✓ 839ms | ✓ 684ms | http |
| 38.34.179.61:8445 | ✓ 833ms | ✓ 1484ms | ✓ 577ms | ✓ 835ms | ✓ 658ms | http |
| 38.34.179.98:8453 | ✓ 837ms | ✓ 1154ms | ✓ 927ms | ✓ 866ms | ✓ 641ms | http |
| 38.34.179.27:8451 | ✓ 797ms | ✓ 1384ms | ✓ 697ms | ✓ 904ms | ✓ 670ms | http |
| 38.34.179.57:8453 | ✓ 846ms | 否 | ✓ 222ms | ✓ 870ms | ✓ 685ms | http |
| 38.34.179.26:8450 | ✓ 809ms | 否 | ✓ 266ms | ✓ 897ms | ✓ 659ms | http |
| 38.34.179.6:8449 | ✓ 781ms | ✓ 1631ms | ✓ 282ms | ✓ 984ms | ✓ 743ms | http |
| 38.34.179.75:8453 | ✓ 834ms | ✓ 980ms | ✓ 1124ms | ✓ 1051ms | ✓ 684ms | http |
| 38.34.183.225:8450 | ✓ 838ms | ✓ 1095ms | ✓ 536ms | ✓ 839ms | ✓ 699ms | http |
| 45.136.130.177:8448 | ✓ 821ms | 否 | ✓ 223ms | ✓ 1067ms | ✓ 746ms | http |
| 38.145.208.241:8453 | ✓ 778ms | 否 | ✓ 640ms | ✓ 849ms | ✓ 722ms | http |
| 38.34.179.178:8445 | ✓ 853ms | 否 | ✓ 212ms | ✓ 831ms | ✓ 706ms | http |
| 45.136.131.53:8452 | ✓ 1177ms | ✓ 999ms | ✓ 286ms | ✓ 1022ms | ✓ 1387ms | http |
| 45.136.130.169:8444 | ✓ 814ms | ✓ 1732ms | ✓ 340ms | ✓ 1332ms | ✓ 976ms | http |
| 45.136.130.167:8444 | ✓ 854ms | 否 | ✓ 518ms | ✓ 1154ms | ✓ 1109ms | http |
| 137.220.151.110:6005 | ✓ 1452ms | 否 | ✓ 982ms | ✓ 1350ms | ✓ 965ms | http |
| 137.220.150.22:6005 | ✓ 815ms | 否 | ✓ 1977ms | ✓ 1275ms | ✓ 1033ms | http |
| 113.160.132.26:8080 | ✓ 1715ms | 否 | ✓ 1015ms | ✓ 1503ms | ✓ 1443ms | http |
| 120.92.212.16:7890 | ✓ 1034ms | 否 | 否 | ✓ 1296ms | ✓ 1045ms | http |
| 59.46.216.131:30001 | ✓ 1039ms | ✓ 1407ms | ✓ 1823ms | 否 | ✓ 1416ms | http |
| 8.219.97.248:80 | ✓ 1892ms | 否 | 否 | ✓ 1835ms | ✓ 1553ms | http |
| 167.103.31.122:8800 | 否 | 否 | ✓ 1292ms | ✓ 1963ms | ✓ 1585ms | http |
| 103.113.70.189:1081 | 否 | 否 | ✓ 892ms | ✓ 1071ms | ✓ 786ms | http |
| 38.145.220.8:8452 | ✓ 1646ms | ✓ 981ms | ✓ 402ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 1038ms | ✓ 1321ms | ✓ 837ms | 否 | 否 | http |
| 38.34.179.150:8449 | ✓ 1070ms | ✓ 1279ms | ✓ 942ms | 否 | 否 | http |
| 38.34.183.221:8452 | ✓ 1670ms | ✓ 988ms | ✓ 420ms | 否 | 否 | http |
| 38.34.183.16:8452 | ✓ 1659ms | 否 | ✓ 1865ms | ✓ 1329ms | ✓ 793ms | http |
| 38.145.208.224:8445 | ✓ 219ms | ✓ 800ms | ✓ 205ms | ✓ 964ms | ✓ 834ms | http |
| 38.145.208.193:8452 | ✓ 433ms | ✓ 858ms | ✓ 212ms | ✓ 816ms | ✓ 783ms | http |
| 38.145.203.97:8448 | ✓ 480ms | ✓ 1337ms | ✓ 211ms | ✓ 853ms | ✓ 676ms | http |
| 38.145.203.106:8448 | ✓ 1170ms | ✓ 764ms | ✓ 211ms | ✓ 805ms | ✓ 647ms | http |
| 38.145.208.226:8451 | ✓ 541ms | ✓ 1788ms | ✓ 381ms | ✓ 1131ms | ✓ 1093ms | http |
| 106.14.203.63:3333 | ✓ 961ms | ✓ 1105ms | ✓ 952ms | ✓ 1165ms | ✓ 922ms | http |
| 137.220.150.152:6005 | 否 | ✓ 1987ms | ✓ 929ms | ✓ 1210ms | ✓ 1032ms | http |
| 147.161.239.240:8800 | ✓ 1093ms | ✓ 1650ms | ✓ 1567ms | ✓ 1629ms | ✓ 1502ms | http |
| 91.238.105.64:2024 | ✓ 1123ms | 否 | 否 | ✓ 1982ms | ✓ 1669ms | http |
| 101.43.127.100:8877 | ✓ 1353ms | 否 | ✓ 1585ms | ✓ 1792ms | 否 | http |
| 137.220.150.170:6005 | ✓ 1418ms | 否 | ✓ 1234ms | ✓ 1248ms | ✓ 1054ms | http |
| 45.93.29.147:6005 | ✓ 1783ms | 否 | 否 | ✓ 1432ms | ✓ 1386ms | http |
| 38.145.220.20:8444 | ✓ 795ms | ✓ 1424ms | ✓ 318ms | ✓ 1108ms | ✓ 1557ms | http |
| 120.92.212.16:8890 | ✓ 1705ms | ✓ 1522ms | ✓ 1279ms | ✓ 1553ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1099ms | 否 | 否 | ✓ 1182ms | ✓ 1169ms | http |
| 34.141.27.50:3128 | ✓ 967ms | ✓ 1461ms | 否 | ✓ 1691ms | 否 | http |
| 142.171.224.229:7890 | ✓ 578ms | 否 | ✓ 961ms | ✓ 786ms | ✓ 596ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 616ms | ✓ 946ms | ✓ 733ms | http |
| 103.184.99.194:8080 | ✓ 1512ms | 否 | ✓ 1992ms | ✓ 1620ms | ✓ 1555ms | http |
| 88.80.150.82:8080 | ✓ 1687ms | 否 | 否 | ✓ 1979ms | ✓ 1922ms | https |
| 114.237.77.244:1080 | ✓ 1952ms | ✓ 1342ms | ✓ 943ms | ✓ 1312ms | ✓ 1773ms | http |
| 103.84.95.54:7890 | ✓ 880ms | 否 | ✓ 1436ms | ✓ 909ms | 否 | http |
| 38.145.218.102:8444 | ✓ 924ms | ✓ 1163ms | ✓ 235ms | ✓ 980ms | ✓ 657ms | http |
| 38.145.208.192:8448 | ✓ 349ms | ✓ 1280ms | ✓ 202ms | ✓ 779ms | ✓ 617ms | http |
| 38.34.183.13:8449 | ✓ 484ms | ✓ 1433ms | ✓ 432ms | ✓ 1575ms | ✓ 874ms | http |
| 38.145.208.189:8446 | ✓ 763ms | ✓ 1375ms | ✓ 194ms | ✓ 803ms | ✓ 1235ms | http |
| 38.145.208.197:8451 | ✓ 526ms | 否 | ✓ 206ms | ✓ 904ms | ✓ 1770ms | http |
| 45.136.131.43:8444 | ✓ 216ms | ✓ 1408ms | ✓ 635ms | ✓ 1050ms | ✓ 794ms | http |
| 38.145.208.253:8445 | ✓ 942ms | ✓ 810ms | ✓ 809ms | 否 | ✓ 656ms | http |
| 38.145.203.135:8451 | ✓ 906ms | 否 | ✓ 754ms | ✓ 1586ms | ✓ 803ms | http |
| 45.136.131.66:8445 | ✓ 1095ms | ✓ 1829ms | ✓ 910ms | ✓ 1098ms | ✓ 1738ms | http |
| 38.145.208.215:8453 | ✓ 244ms | ✓ 1187ms | ✓ 327ms | ✓ 1016ms | ✓ 1481ms | http |
| 38.34.179.160:8453 | ✓ 271ms | ✓ 1878ms | ✓ 735ms | ✓ 1624ms | ✓ 638ms | http |
| 38.34.179.155:8452 | ✓ 515ms | ✓ 1341ms | ✓ 1992ms | ✓ 1122ms | ✓ 625ms | http |
| 38.34.179.160:8452 | ✓ 511ms | ✓ 1336ms | ✓ 1997ms | ✓ 1125ms | ✓ 628ms | http |
| 45.136.131.29:8450 | ✓ 576ms | ✓ 1553ms | 否 | ✓ 1250ms | 否 | http |
| 38.145.220.72:8444 | ✓ 489ms | ✓ 1435ms | ✓ 665ms | ✓ 822ms | ✓ 812ms | http |
| 38.145.220.56:8453 | ✓ 488ms | 否 | ✓ 194ms | ✓ 799ms | ✓ 756ms | http |
| 38.145.208.194:8453 | ✓ 1602ms | ✓ 908ms | ✓ 304ms | ✓ 1352ms | ✓ 1929ms | http |
| 38.145.208.188:8447 | ✓ 1352ms | ✓ 1800ms | ✓ 939ms | 否 | 否 | http |
| 45.136.131.32:8451 | ✓ 1805ms | ✓ 909ms | ✓ 804ms | ✓ 1133ms | ✓ 700ms | http |
| 45.136.131.36:8450 | ✓ 1779ms | ✓ 1378ms | ✓ 203ms | ✓ 885ms | ✓ 621ms | http |
| 45.136.131.31:8451 | ✓ 1807ms | ✓ 975ms | ✓ 952ms | ✓ 1049ms | ✓ 725ms | http |
| 45.136.131.34:8451 | ✓ 1778ms | 否 | ✓ 228ms | ✓ 986ms | ✓ 798ms | http |
| 45.136.131.30:8451 | ✓ 1798ms | 否 | ✓ 228ms | ✓ 1012ms | ✓ 775ms | http |
| 38.34.179.162:8451 | ✓ 1822ms | ✓ 1238ms | ✓ 434ms | ✓ 1125ms | ✓ 731ms | http |
| 45.136.131.25:8450 | ✓ 1784ms | ✓ 1278ms | ✓ 822ms | ✓ 941ms | ✓ 746ms | http |
| 38.34.179.106:8446 | ✓ 1808ms | ✓ 1625ms | ✓ 727ms | ✓ 809ms | ✓ 1049ms | http |
| 193.23.200.251:10808 | ✓ 1091ms | 否 | ✓ 857ms | 否 | ✓ 1544ms | http |
| 38.34.179.151:8452 | ✓ 1794ms | ✓ 1432ms | ✓ 1093ms | ✓ 824ms | ✓ 1527ms | http |
| 38.34.179.20:8445 | ✓ 1798ms | 否 | ✓ 1468ms | ✓ 1406ms | 否 | http |

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
