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

最后更新：2026-03-19 12:27:46 UTC（2026-03-19 20:27:46 UTC+8）

**代理总数：88**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 664ms | ✓ 1084ms | ✓ 944ms | ✓ 1069ms | ✓ 814ms | http |
| 202.155.12.161:443 | ✓ 1009ms | 否 | ✓ 927ms | ✓ 1351ms | 否 | http |
| 219.117.204.211:7799 | ✓ 1218ms | 否 | ✓ 483ms | ✓ 844ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1623ms | 否 | ✓ 1297ms | ✓ 1185ms | ✓ 1353ms | http |
| 137.220.150.170:6005 | ✓ 1614ms | 否 | ✓ 1013ms | ✓ 1224ms | ✓ 933ms | http |
| 85.198.96.242:3128 | ✓ 749ms | 否 | ✓ 708ms | 否 | ✓ 1420ms | http |
| 101.47.73.135:3128 | ✓ 918ms | 否 | ✓ 1807ms | ✓ 979ms | 否 | http |
| 8.219.97.248:80 | ✓ 1596ms | 否 | ✓ 934ms | ✓ 1473ms | 否 | http |
| 45.125.67.37:443 | ✓ 1612ms | 否 | ✓ 1083ms | ✓ 1139ms | ✓ 1039ms | http |
| 147.161.239.240:8800 | ✓ 1849ms | ✓ 1832ms | ✓ 1189ms | ✓ 1474ms | ✓ 1258ms | http |
| 137.220.150.104:6005 | ✓ 1676ms | 否 | ✓ 768ms | ✓ 1098ms | ✓ 877ms | http |
| 222.184.48.251:22222 | ✓ 1701ms | ✓ 1207ms | ✓ 1378ms | 否 | ✓ 1637ms | http |
| 45.149.92.147:5001 | ✓ 742ms | 否 | ✓ 982ms | ✓ 869ms | ✓ 647ms | http |
| 137.220.151.110:6005 | ✓ 851ms | 否 | ✓ 1031ms | ✓ 1159ms | ✓ 895ms | http |
| 138.124.53.25:7443 | ✓ 1220ms | 否 | ✓ 978ms | 否 | ✓ 1654ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1124ms | 否 | ✓ 1003ms | ✓ 1002ms | http |
| 38.55.106.206:6005 | ✓ 1021ms | 否 | ✓ 1087ms | ✓ 864ms | ✓ 1040ms | http |
| 38.34.179.156:8450 | ✓ 1161ms | 否 | ✓ 1372ms | ✓ 698ms | ✓ 515ms | http |
| 101.43.127.100:8877 | ✓ 924ms | ✓ 1130ms | ✓ 1748ms | ✓ 1305ms | ✓ 1109ms | http |
| 45.136.130.173:8448 | ✓ 1471ms | 否 | ✓ 1510ms | ✓ 1932ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1571ms | 否 | ✓ 1137ms | 否 | ✓ 1221ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1454ms | 否 | ✓ 1945ms | ✓ 1967ms | http |
| 35.225.22.61:80 | ✓ 437ms | 否 | ✓ 1058ms | ✓ 1271ms | ✓ 1022ms | http |
| 178.236.245.17:3128 | ✓ 1253ms | 否 | ✓ 819ms | 否 | ✓ 1477ms | http |
| 34.101.184.164:3128 | ✓ 1641ms | 否 | 否 | ✓ 1293ms | ✓ 1107ms | http |
| 183.249.5.111:22222 | 否 | ✓ 918ms | ✓ 724ms | ✓ 1205ms | ✓ 1925ms | http |
| 185.191.236.162:3128 | ✓ 1067ms | ✓ 1810ms | 否 | ✓ 1998ms | ✓ 1280ms | http |
| 143.244.140.119:3128 | ✓ 1443ms | 否 | 否 | ✓ 1867ms | ✓ 1644ms | http |
| 38.55.107.254:6005 | ✓ 1459ms | 否 | ✓ 1316ms | ✓ 1197ms | ✓ 1684ms | http |
| 59.46.216.131:30001 | ✓ 999ms | 否 | ✓ 1102ms | ✓ 1292ms | ✓ 1117ms | http |
| 174.138.24.77:1080 | ✓ 1576ms | 否 | ✓ 1684ms | ✓ 1600ms | 否 | http |
| 183.249.5.117:22222 | ✓ 1894ms | ✓ 861ms | ✓ 732ms | ✓ 1213ms | ✓ 817ms | http |
| 162.243.149.86:31028 | ✓ 890ms | 否 | ✓ 1090ms | ✓ 801ms | ✓ 632ms | http |
| 168.235.110.63:3128 | ✓ 832ms | ✓ 1241ms | ✓ 1661ms | ✓ 1467ms | ✓ 1009ms | http |
| 103.113.70.189:1081 | ✓ 834ms | 否 | ✓ 1073ms | ✓ 1460ms | ✓ 1038ms | http |
| 183.249.5.105:22222 | 否 | ✓ 1925ms | ✓ 724ms | ✓ 895ms | ✓ 1986ms | http |
| 38.34.179.61:8445 | 否 | ✓ 847ms | ✓ 718ms | ✓ 1966ms | ✓ 992ms | http |
| 45.136.131.26:8451 | ✓ 372ms | 否 | ✓ 1404ms | ✓ 687ms | ✓ 913ms | http |
| 38.55.106.208:6005 | ✓ 1327ms | 否 | ✓ 1038ms | ✓ 1157ms | ✓ 747ms | http |
| 164.92.148.68:3128 | ✓ 578ms | 否 | 否 | ✓ 1743ms | ✓ 1419ms | http |
| 38.145.208.243:8445 | ✓ 385ms | ✓ 1815ms | ✓ 1838ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1345ms | 否 | ✓ 1500ms | ✓ 983ms | ✓ 1035ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1502ms | ✓ 1134ms | ✓ 994ms | http |
| 194.5.212.40:8080 | ✓ 1335ms | 否 | ✓ 1477ms | 否 | ✓ 1754ms | http |
| 116.80.96.102:3172 | ✓ 1514ms | 否 | ✓ 1484ms | 否 | ✓ 1804ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 962ms | ✓ 1699ms | ✓ 1267ms | http |
| 106.117.208.101:7890 | ✓ 1566ms | ✓ 1405ms | ✓ 1471ms | ✓ 1494ms | 否 | http |
| 38.145.208.244:8448 | ✓ 1769ms | 否 | ✓ 346ms | ✓ 703ms | ✓ 811ms | http |
| 38.145.208.240:8448 | ✓ 1748ms | 否 | ✓ 352ms | ✓ 737ms | ✓ 813ms | http |
| 120.92.212.16:8890 | ✓ 966ms | 否 | ✓ 978ms | ✓ 1253ms | ✓ 973ms | http |
| 120.92.212.16:7890 | ✓ 971ms | 否 | ✓ 982ms | ✓ 1319ms | 否 | http |
| 121.40.231.103:7890 | ✓ 900ms | ✓ 1000ms | ✓ 852ms | 否 | ✓ 1329ms | http |
| 180.125.216.109:8118 | ✓ 1877ms | 否 | ✓ 1038ms | 否 | ✓ 1629ms | http |
| 183.249.5.110:22222 | 否 | ✓ 1138ms | ✓ 859ms | 否 | ✓ 1032ms | http |
| 103.82.23.118:5247 | 否 | 否 | ✓ 1262ms | ✓ 1864ms | ✓ 1786ms | http |
| 114.237.77.231:1080 | ✓ 1960ms | ✓ 1739ms | ✓ 930ms | ✓ 1307ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1395ms | 否 | ✓ 1826ms | 否 | ✓ 1975ms | http |
| 45.136.130.171:8445 | ✓ 1523ms | ✓ 784ms | 否 | ✓ 1294ms | ✓ 1810ms | http |
| 38.55.107.137:6005 | ✓ 1823ms | 否 | ✓ 1304ms | ✓ 1117ms | ✓ 1168ms | http |
| 38.145.203.34:8445 | ✓ 965ms | ✓ 737ms | ✓ 693ms | ✓ 1805ms | ✓ 571ms | http |
| 165.225.72.38:11178 | ✓ 1901ms | ✓ 1729ms | ✓ 1114ms | 否 | 否 | http |
| 165.225.72.38:11231 | ✓ 1208ms | 否 | ✓ 943ms | ✓ 1678ms | ✓ 1309ms | http |
| 165.225.72.38:10975 | ✓ 1615ms | 否 | ✓ 624ms | ✓ 1758ms | ✓ 1139ms | http |
| 165.225.72.38:10300 | ✓ 1613ms | 否 | ✓ 624ms | ✓ 1644ms | ✓ 1254ms | http |
| 165.225.72.38:12000 | ✓ 1616ms | 否 | ✓ 621ms | ✓ 1701ms | ✓ 1228ms | http |
| 165.225.72.38:10319 | 否 | ✓ 1625ms | ✓ 607ms | ✓ 1646ms | ✓ 1323ms | http |
| 165.225.72.38:11258 | ✓ 1210ms | 否 | ✓ 1027ms | ✓ 1717ms | ✓ 1295ms | http |
| 165.225.72.38:10974 | 否 | 否 | ✓ 609ms | ✓ 1450ms | ✓ 1169ms | http |
| 165.225.72.38:11404 | 否 | 否 | ✓ 600ms | ✓ 1466ms | ✓ 1163ms | http |
| 165.225.72.38:11350 | 否 | ✓ 1645ms | ✓ 587ms | ✓ 1688ms | ✓ 1310ms | http |
| 165.225.72.38:10173 | 否 | ✓ 1703ms | ✓ 715ms | ✓ 1455ms | ✓ 1291ms | http |
| 165.225.72.38:10418 | 否 | 否 | ✓ 610ms | ✓ 1480ms | ✓ 1182ms | http |
| 165.225.72.38:11169 | ✓ 1668ms | 否 | ✓ 667ms | ✓ 1612ms | ✓ 1304ms | http |
| 165.225.72.38:9443 | 否 | 否 | ✓ 639ms | ✓ 1524ms | ✓ 1148ms | http |
| 165.225.72.38:10676 | ✓ 1686ms | ✓ 1672ms | ✓ 1571ms | ✓ 1464ms | ✓ 1166ms | http |
| 165.225.72.38:11310 | ✓ 1616ms | 否 | ✓ 622ms | ✓ 1616ms | ✓ 1247ms | http |
| 165.225.72.38:10922 | 否 | ✓ 1713ms | ✓ 616ms | ✓ 1621ms | ✓ 1227ms | http |
| 165.225.72.38:10637 | ✓ 1614ms | 否 | ✓ 623ms | ✓ 1739ms | ✓ 1122ms | http |
| 38.34.179.176:8450 | ✓ 554ms | 否 | ✓ 895ms | ✓ 723ms | ✓ 682ms | http |
| 38.145.203.111:8444 | ✓ 555ms | 否 | ✓ 1740ms | ✓ 1130ms | 否 | http |
| 1.234.153.14:80 | ✓ 1493ms | 否 | ✓ 758ms | ✓ 1583ms | 否 | http |
| 103.183.10.169:3125 | ✓ 1386ms | 否 | 否 | ✓ 1590ms | ✓ 1476ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1892ms | ✓ 1545ms | ✓ 1417ms | http |
| 76.72.147.141:3128 | ✓ 1792ms | ✓ 1446ms | ✓ 1820ms | 否 | 否 | http |
| 165.225.72.38:10539 | ✓ 1054ms | ✓ 1711ms | 否 | 否 | ✓ 1131ms | http |
| 165.225.72.38:11447 | ✓ 1056ms | 否 | ✓ 591ms | ✓ 1471ms | 否 | http |
| 133.242.138.34:8100 | 否 | 否 | ✓ 616ms | ✓ 1402ms | ✓ 744ms | http |
| 120.55.163.237:10086 | 否 | ✓ 1163ms | ✓ 989ms | 否 | ✓ 922ms | http |

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
