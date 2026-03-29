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

最后更新：2026-03-29 17:36:29 UTC（2026-03-30 01:36:29 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 87 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 207ms | ✓ 722ms | ✓ 850ms | ✓ 697ms | ✓ 530ms | http |
| 39.185.46.193:5911 | ✓ 632ms | ✓ 792ms | ✓ 804ms | ✓ 868ms | ✓ 663ms | http |
| 1.231.81.166:3128 | ✓ 1313ms | ✓ 968ms | ✓ 1768ms | ✓ 932ms | ✓ 847ms | http |
| 147.161.210.140:8800 | ✓ 1282ms | 否 | ✓ 1524ms | ✓ 913ms | ✓ 923ms | http |
| 219.117.204.211:7799 | ✓ 1284ms | 否 | 否 | ✓ 791ms | ✓ 607ms | http |
| 42.96.16.158:1311 | ✓ 1352ms | 否 | ✓ 941ms | ✓ 1124ms | ✓ 890ms | http |
| 147.161.239.240:8800 | ✓ 1379ms | ✓ 1784ms | ✓ 1140ms | ✓ 1494ms | ✓ 1692ms | http |
| 167.103.115.102:8800 | ✓ 1049ms | 否 | ✓ 1165ms | ✓ 1214ms | 否 | http |
| 167.103.34.108:8800 | ✓ 1395ms | 否 | ✓ 1299ms | ✓ 1287ms | ✓ 1227ms | http |
| 95.213.217.168:52004 | ✓ 1402ms | 否 | ✓ 1265ms | ✓ 1837ms | ✓ 1362ms | http |
| 113.160.132.26:8080 | ✓ 1360ms | 否 | ✓ 1091ms | 否 | ✓ 1404ms | http |
| 35.225.22.61:80 | 否 | ✓ 1365ms | ✓ 429ms | ✓ 1081ms | ✓ 945ms | http |
| 43.99.54.236:5555 | ✓ 665ms | ✓ 1172ms | ✓ 649ms | ✓ 771ms | 否 | http |
| 183.249.5.117:22222 | ✓ 709ms | ✓ 829ms | ✓ 659ms | ✓ 1129ms | ✓ 677ms | http |
| 222.184.48.251:22222 | ✓ 807ms | ✓ 1748ms | ✓ 964ms | ✓ 1175ms | ✓ 823ms | http |
| 167.103.144.127:8800 | ✓ 1575ms | 否 | ✓ 1247ms | ✓ 1541ms | ✓ 1334ms | http |
| 103.84.95.54:7890 | ✓ 754ms | 否 | ✓ 1958ms | ✓ 1173ms | 否 | http |
| 222.184.48.242:22222 | 否 | ✓ 1091ms | ✓ 971ms | ✓ 1770ms | ✓ 797ms | http |
| 101.47.73.135:3128 | ✓ 1851ms | 否 | ✓ 1720ms | ✓ 1912ms | 否 | http |
| 45.144.28.81:10808 | ✓ 1255ms | 否 | ✓ 836ms | ✓ 1668ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1317ms | 否 | ✓ 1297ms | ✓ 1622ms | 否 | http |
| 59.46.216.131:30001 | ✓ 958ms | 否 | ✓ 1070ms | ✓ 1247ms | ✓ 1047ms | http |
| 45.12.151.226:2829 | 否 | 否 | ✓ 1655ms | ✓ 1956ms | ✓ 1653ms | http |
| 1.234.153.14:80 | ✓ 1339ms | ✓ 937ms | ✓ 725ms | ✓ 902ms | ✓ 716ms | http |
| 101.43.127.100:8877 | ✓ 900ms | ✓ 1052ms | ✓ 1109ms | ✓ 1065ms | ✓ 857ms | http |
| 177.234.217.88:999 | ✓ 1471ms | 否 | ✓ 1940ms | 否 | ✓ 1719ms | http |
| 45.136.131.40:8444 | ✓ 157ms | 否 | ✓ 84ms | ✓ 686ms | ✓ 684ms | http |
| 38.34.183.225:8444 | ✓ 329ms | ✓ 1165ms | ✓ 101ms | ✓ 705ms | ✓ 638ms | http |
| 38.34.179.40:8446 | ✓ 516ms | ✓ 709ms | ✓ 214ms | ✓ 988ms | 否 | http |
| 45.136.131.61:8451 | ✓ 846ms | ✓ 640ms | ✓ 115ms | ✓ 820ms | ✓ 930ms | http |
| 38.34.179.62:8450 | ✓ 1451ms | ✓ 1408ms | ✓ 340ms | ✓ 988ms | ✓ 1419ms | http |
| 59.8.203.55:80 | ✓ 695ms | 否 | ✓ 946ms | ✓ 1282ms | ✓ 741ms | http |
| 38.34.179.79:8451 | ✓ 271ms | ✓ 1688ms | 否 | ✓ 665ms | ✓ 897ms | http |
| 152.69.229.220:3128 | ✓ 801ms | ✓ 1607ms | ✓ 1078ms | ✓ 1179ms | ✓ 1677ms | http |
| 5.104.87.17:8051 | ✓ 1448ms | 否 | ✓ 1913ms | ✓ 1958ms | ✓ 1493ms | http |
| 38.145.218.27:8446 | ✓ 1511ms | 否 | ✓ 86ms | ✓ 670ms | 否 | http |
| 45.136.130.183:8453 | ✓ 851ms | 否 | ✓ 1922ms | ✓ 1607ms | 否 | http |
| 183.249.5.110:22222 | ✓ 723ms | ✓ 846ms | ✓ 711ms | ✓ 923ms | ✓ 679ms | http |
| 183.249.5.111:22222 | ✓ 714ms | ✓ 1047ms | ✓ 909ms | ✓ 909ms | ✓ 686ms | http |
| 120.92.212.16:8890 | ✓ 1243ms | 否 | ✓ 1677ms | ✓ 1814ms | ✓ 1184ms | http |
| 121.230.9.225:1080 | 否 | 否 | ✓ 1225ms | ✓ 1352ms | ✓ 1063ms | http |
| 120.92.212.16:7890 | ✓ 1757ms | ✓ 1901ms | ✓ 947ms | ✓ 1174ms | ✓ 1149ms | http |
| 121.230.8.213:1080 | ✓ 1679ms | 否 | ✓ 965ms | ✓ 1394ms | ✓ 1046ms | http |
| 20.210.76.178:8561 | ✓ 691ms | ✓ 1011ms | ✓ 1104ms | ✓ 1478ms | ✓ 1585ms | http |
| 193.233.22.29:10808 | ✓ 1706ms | 否 | ✓ 1804ms | ✓ 1484ms | 否 | http |
| 20.210.76.175:8561 | ✓ 733ms | ✓ 1247ms | ✓ 628ms | ✓ 947ms | ✓ 810ms | http |
| 20.210.76.104:8561 | ✓ 730ms | ✓ 1217ms | ✓ 633ms | ✓ 973ms | ✓ 811ms | http |
| 222.228.171.92:8080 | ✓ 1494ms | 否 | ✓ 1317ms | ✓ 1149ms | ✓ 844ms | http |
| 8.222.175.80:6128 | ✓ 698ms | ✓ 1612ms | ✓ 683ms | ✓ 990ms | ✓ 786ms | http |
| 106.75.15.167:7890 | ✓ 1349ms | 否 | 否 | ✓ 1213ms | ✓ 1731ms | http |
| 5.102.109.41:999 | ✓ 843ms | ✓ 1340ms | ✓ 1835ms | ✓ 1543ms | ✓ 1421ms | http |
| 116.80.96.103:3172 | ✓ 1732ms | 否 | 否 | ✓ 1802ms | ✓ 1775ms | http |
| 183.249.5.105:22222 | ✓ 694ms | ✓ 839ms | 否 | ✓ 992ms | 否 | http |
| 38.34.179.21:8446 | ✓ 1516ms | ✓ 1103ms | ✓ 141ms | ✓ 1044ms | 否 | http |
| 45.136.131.29:8453 | ✓ 1758ms | 否 | ✓ 997ms | ✓ 1697ms | ✓ 581ms | http |
| 34.101.184.164:3128 | ✓ 1820ms | 否 | ✓ 1186ms | ✓ 1933ms | ✓ 1588ms | http |
| 157.230.220.25:4857 | ✓ 428ms | 否 | 否 | ✓ 1457ms | ✓ 1244ms | http |
| 20.27.15.49:8561 | ✓ 447ms | ✓ 912ms | ✓ 435ms | ✓ 726ms | ✓ 570ms | http |
| 38.145.208.169:8446 | ✓ 1490ms | ✓ 739ms | ✓ 307ms | ✓ 715ms | ✓ 774ms | http |
| 38.145.208.171:8453 | ✓ 1801ms | ✓ 1260ms | ✓ 1975ms | ✓ 1361ms | ✓ 777ms | http |
| 38.34.179.29:8449 | ✓ 555ms | ✓ 911ms | ✓ 1565ms | ✓ 1665ms | ✓ 692ms | http |
| 38.34.179.100:8452 | ✓ 712ms | ✓ 1434ms | 否 | ✓ 1211ms | ✓ 587ms | http |
| 38.34.179.66:8446 | ✓ 831ms | 否 | ✓ 1241ms | ✓ 928ms | ✓ 1656ms | http |
| 168.110.52.228:3128 | 否 | 否 | ✓ 1515ms | ✓ 1923ms | ✓ 973ms | http |
| 85.208.108.43:2094 | ✓ 620ms | 否 | ✓ 1228ms | ✓ 1391ms | ✓ 1020ms | http |
| 106.117.208.101:7890 | 否 | 否 | ✓ 945ms | ✓ 1159ms | ✓ 985ms | http |
| 45.140.147.155:1081 | ✓ 1079ms | 否 | ✓ 1048ms | ✓ 1482ms | ✓ 1225ms | http |
| 8.219.97.248:80 | ✓ 1555ms | 否 | ✓ 1793ms | ✓ 1622ms | 否 | http |
| 118.31.1.154:80 | ✓ 799ms | ✓ 1070ms | ✓ 867ms | ✓ 1223ms | ✓ 850ms | http |
| 84.38.185.139:3128 | ✓ 1422ms | 否 | ✓ 1468ms | 否 | ✓ 1771ms | http |
| 45.136.198.40:3128 | ✓ 1434ms | ✓ 1895ms | ✓ 1884ms | 否 | ✓ 1924ms | http |
| 198.59.68.130:3128 | ✓ 988ms | 否 | 否 | ✓ 1430ms | ✓ 1332ms | http |
| 175.194.173.105:3128 | ✓ 1916ms | ✓ 1030ms | ✓ 1882ms | 否 | 否 | http |
| 20.2.83.243:3128 | ✓ 670ms | ✓ 1891ms | ✓ 1051ms | ✓ 1279ms | ✓ 715ms | http |
| 20.78.213.56:80 | ✓ 1959ms | ✓ 1175ms | 否 | ✓ 1079ms | ✓ 763ms | http |
| 38.34.179.164:8448 | ✓ 339ms | ✓ 1784ms | ✓ 1185ms | ✓ 693ms | ✓ 662ms | http |
| 124.121.2.189:8080 | ✓ 1569ms | 否 | 否 | ✓ 1522ms | ✓ 1466ms | http |
| 38.145.208.240:8452 | ✓ 869ms | ✓ 1662ms | ✓ 1044ms | ✓ 805ms | ✓ 1634ms | http |
| 45.136.130.171:8445 | ✓ 1080ms | ✓ 1130ms | ✓ 641ms | 否 | ✓ 932ms | http |
| 38.34.183.221:8445 | ✓ 189ms | ✓ 603ms | ✓ 498ms | ✓ 1819ms | ✓ 1542ms | http |
| 103.52.115.171:3128 | ✓ 1570ms | 否 | ✓ 773ms | ✓ 1149ms | ✓ 967ms | http |
| 149.62.191.202:3128 | ✓ 1310ms | 否 | ✓ 1884ms | 否 | ✓ 1433ms | http |
| 103.39.51.190:8080 | ✓ 1500ms | 否 | 否 | ✓ 1500ms | ✓ 1761ms | http |
| 38.34.178.175:8444 | ✓ 192ms | ✓ 606ms | ✓ 208ms | ✓ 905ms | ✓ 1664ms | http |
| 38.145.208.247:8452 | 否 | ✓ 965ms | ✓ 1079ms | 否 | ✓ 583ms | http |
| 38.145.208.243:8452 | 否 | ✓ 1343ms | ✓ 1032ms | 否 | ✓ 523ms | http |
| 45.140.147.82:1081 | ✓ 862ms | ✓ 1537ms | ✓ 1507ms | 否 | ✓ 1496ms | http |
| 88.80.150.82:8080 | ✓ 1141ms | ✓ 1936ms | 否 | 否 | ✓ 1880ms | https |

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
