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

最后更新：2026-03-07 21:21:32 UTC（2026-03-08 05:21:32 UTC+8）

**代理总数：81**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 546ms | ✓ 962ms | ✓ 816ms | ✓ 1127ms | ✓ 902ms | http |
| 1.231.81.166:3128 | ✓ 1833ms | ✓ 1067ms | ✓ 1142ms | ✓ 1052ms | ✓ 1024ms | http |
| 67.169.98.211:443 | ✓ 1186ms | 否 | 否 | ✓ 1766ms | ✓ 980ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1373ms | ✓ 1076ms | ✓ 1370ms | 否 | http |
| 217.76.245.80:999 | ✓ 849ms | ✓ 1447ms | ✓ 1223ms | ✓ 1542ms | ✓ 1233ms | http |
| 85.9.195.140:1080 | ✓ 735ms | 否 | ✓ 1192ms | ✓ 1349ms | ✓ 1046ms | http |
| 115.231.181.40:8128 | ✓ 1984ms | ✓ 1421ms | ✓ 1062ms | ✓ 1322ms | 否 | http |
| 116.80.82.225:3172 | ✓ 1665ms | 否 | ✓ 1675ms | ✓ 1984ms | ✓ 1783ms | http |
| 193.228.139.78:8888 | ✓ 576ms | 否 | ✓ 1322ms | 否 | ✓ 1507ms | http |
| 45.140.147.82:1081 | ✓ 555ms | ✓ 1210ms | 否 | 否 | ✓ 1214ms | http |
| 62.113.119.14:8080 | ✓ 651ms | 否 | ✓ 1314ms | ✓ 1938ms | ✓ 1483ms | http |
| 211.171.114.154:3128 | ✓ 1833ms | 否 | ✓ 1844ms | ✓ 1398ms | 否 | http |
| 152.42.213.210:8080 | ✓ 1513ms | 否 | ✓ 1563ms | ✓ 1267ms | 否 | http |
| 193.168.173.136:443 | ✓ 765ms | 否 | ✓ 1315ms | 否 | ✓ 1786ms | http |
| 101.43.255.96:80 | ✓ 1113ms | ✓ 1380ms | ✓ 1188ms | ✓ 1333ms | ✓ 1077ms | http |
| 81.70.169.194:80 | ✓ 1135ms | ✓ 1395ms | ✓ 1111ms | ✓ 1332ms | ✓ 1092ms | http |
| 14.56.107.244:3128 | ✓ 1618ms | 否 | ✓ 1149ms | 否 | ✓ 869ms | http |
| 121.128.121.54:3128 | ✓ 1621ms | 否 | 否 | ✓ 1112ms | ✓ 1920ms | http |
| 45.186.6.104:3128 | ✓ 1260ms | ✓ 1848ms | ✓ 1664ms | 否 | 否 | http |
| 42.115.72.27:2033 | ✓ 1632ms | 否 | ✓ 1688ms | ✓ 1904ms | 否 | http |
| 42.115.72.27:2038 | ✓ 1698ms | 否 | ✓ 1606ms | ✓ 1910ms | ✓ 1690ms | http |
| 125.128.12.144:3128 | ✓ 1827ms | ✓ 1987ms | ✓ 1516ms | ✓ 1934ms | 否 | http |
| 39.104.201.40:7890 | ✓ 1012ms | ✓ 1325ms | ✓ 1117ms | ✓ 1365ms | 否 | http |
| 190.9.109.205:999 | ✓ 881ms | ✓ 1196ms | ✓ 1143ms | ✓ 1474ms | ✓ 1124ms | http |
| 190.9.109.207:999 | ✓ 955ms | ✓ 1154ms | ✓ 1101ms | ✓ 1497ms | ✓ 1026ms | http |
| 187.102.219.64:999 | ✓ 1263ms | 否 | ✓ 1640ms | 否 | ✓ 1943ms | http |
| 27.73.57.47:10005 | ✓ 1494ms | 否 | ✓ 1893ms | ✓ 1766ms | ✓ 1783ms | http |
| 138.124.53.25:7443 | ✓ 1683ms | 否 | ✓ 1772ms | 否 | ✓ 1692ms | http |
| 61.72.110.54:3128 | 否 | ✓ 1639ms | ✓ 1496ms | ✓ 1439ms | ✓ 1210ms | http |
| 35.225.22.61:80 | ✓ 735ms | 否 | ✓ 176ms | 否 | ✓ 650ms | http |
| 209.38.51.97:3128 | 否 | ✓ 1927ms | ✓ 203ms | 否 | ✓ 968ms | http |
| 165.227.5.10:8888 | 否 | ✓ 823ms | ✓ 603ms | ✓ 1181ms | ✓ 689ms | http |
| 202.155.12.161:443 | ✓ 1920ms | 否 | 否 | ✓ 1484ms | ✓ 1115ms | http |
| 14.225.217.30:7890 | 否 | 否 | ✓ 992ms | ✓ 1265ms | ✓ 1067ms | http |
| 178.236.245.17:3128 | ✓ 922ms | 否 | ✓ 1599ms | ✓ 1974ms | ✓ 1544ms | http |
| 178.236.245.59:3128 | ✓ 555ms | 否 | ✓ 777ms | ✓ 1632ms | ✓ 1368ms | http |
| 61.72.221.94:3128 | 否 | 否 | ✓ 719ms | ✓ 1206ms | ✓ 932ms | http |
| 61.72.221.194:3128 | ✓ 1824ms | ✓ 1618ms | ✓ 1363ms | ✓ 1370ms | 否 | http |
| 14.225.222.164:7890 | ✓ 1515ms | 否 | 否 | ✓ 1515ms | ✓ 1657ms | http |
| 94.131.122.35:9000 | ✓ 620ms | 否 | ✓ 1534ms | ✓ 1856ms | ✓ 1821ms | http |
| 120.92.212.16:7890 | ✓ 1120ms | ✓ 1344ms | 否 | ✓ 1331ms | ✓ 1807ms | http |
| 138.197.68.35:4857 | 否 | ✓ 1614ms | ✓ 480ms | ✓ 1904ms | 否 | http |
| 45.140.147.82:1082 | ✓ 406ms | 否 | ✓ 652ms | ✓ 1262ms | ✓ 938ms | http |
| 15.204.233.75:3128 | 否 | 否 | ✓ 1896ms | ✓ 1367ms | ✓ 1062ms | http |
| 192.166.82.55:1080 | ✓ 1176ms | ✓ 1935ms | ✓ 1097ms | ✓ 1282ms | ✓ 1538ms | http |
| 162.248.165.72:1080 | ✓ 848ms | 否 | ✓ 639ms | ✓ 1949ms | ✓ 1585ms | http |
| 121.230.9.64:1080 | ✓ 1255ms | ✓ 1596ms | ✓ 1599ms | 否 | 否 | http |
| 103.215.36.88:18989 | ✓ 1214ms | ✓ 1394ms | ✓ 1166ms | ✓ 1438ms | ✓ 1099ms | http |
| 185.243.218.43:49153 | ✓ 1610ms | 否 | ✓ 1466ms | ✓ 1970ms | ✓ 1620ms | http |
| 101.32.244.83:8080 | ✓ 1317ms | 否 | ✓ 1059ms | ✓ 1606ms | ✓ 1379ms | http |
| 121.43.196.213:8222 | ✓ 1031ms | ✓ 1206ms | ✓ 1088ms | ✓ 1329ms | ✓ 1022ms | http |
| 121.43.196.210:8222 | ✓ 1038ms | ✓ 1186ms | ✓ 1095ms | ✓ 1309ms | ✓ 1044ms | http |
| 114.55.226.123:10086 | ✓ 1179ms | ✓ 1757ms | ✓ 1124ms | ✓ 1389ms | ✓ 1117ms | http |
| 116.80.82.231:3172 | ✓ 1836ms | 否 | ✓ 1795ms | ✓ 1979ms | ✓ 1821ms | http |
| 168.235.110.63:3128 | ✓ 678ms | ✓ 1235ms | ✓ 135ms | ✓ 1005ms | ✓ 776ms | http |
| 125.128.12.14:3128 | ✓ 1685ms | ✓ 1463ms | ✓ 1438ms | ✓ 1137ms | ✓ 880ms | http |
| 61.76.95.217:40088 | ✓ 1666ms | 否 | ✓ 1285ms | ✓ 1356ms | ✓ 1134ms | http |
| 91.193.240.157:9877 | ✓ 1003ms | 否 | ✓ 1523ms | 否 | ✓ 1739ms | http |
| 61.72.221.234:3128 | ✓ 734ms | ✓ 1597ms | ✓ 1536ms | 否 | ✓ 1055ms | http |
| 103.84.95.54:7890 | ✓ 757ms | 否 | ✓ 786ms | ✓ 1321ms | ✓ 926ms | http |
| 154.64.240.39:1080 | ✓ 1394ms | ✓ 1899ms | ✓ 1348ms | ✓ 1360ms | ✓ 1267ms | http |
| 51.250.37.15:6666 | ✓ 1915ms | ✓ 1910ms | ✓ 1094ms | 否 | ✓ 1898ms | http |
| 46.183.25.8:443 | ✓ 1852ms | 否 | 否 | ✓ 1794ms | ✓ 1337ms | http |
| 220.121.154.240:3128 | ✓ 1880ms | ✓ 1651ms | 否 | ✓ 1125ms | 否 | http |
| 86.109.3.24:9401 | ✓ 431ms | 否 | 否 | ✓ 1926ms | ✓ 842ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1405ms | 否 | ✓ 1176ms | ✓ 983ms | http |
| 172.212.68.37:3128 | ✓ 1582ms | 否 | ✓ 754ms | ✓ 1253ms | ✓ 1054ms | http |
| 45.129.141.143:3128 | ✓ 643ms | 否 | ✓ 1786ms | 否 | ✓ 1694ms | http |
| 89.185.85.138:1080 | ✓ 467ms | ✓ 1844ms | ✓ 990ms | ✓ 1854ms | ✓ 1275ms | http |
| 47.101.159.19:8899 | ✓ 983ms | ✓ 1195ms | 否 | ✓ 1243ms | ✓ 1009ms | http |
| 157.0.142.246:10057 | ✓ 1252ms | ✓ 1372ms | ✓ 1113ms | ✓ 1411ms | ✓ 1074ms | http |
| 103.215.36.88:18113 | ✓ 1153ms | ✓ 1508ms | ✓ 1164ms | ✓ 1540ms | ✓ 1124ms | http |
| 103.139.138.194:3128 | ✓ 1446ms | 否 | ✓ 1828ms | ✓ 1934ms | ✓ 1272ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1371ms | ✓ 1532ms | ✓ 1500ms | http |
| 46.249.103.192:443 | ✓ 1291ms | 否 | ✓ 1256ms | ✓ 1926ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1207ms | ✓ 1515ms | ✓ 1561ms | 否 | ✓ 1616ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 1193ms | ✓ 1190ms | ✓ 983ms | http |
| 86.109.3.24:9443 | ✓ 1243ms | ✓ 1857ms | ✓ 462ms | 否 | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1099ms | ✓ 1617ms | 否 | ✓ 1903ms | http |
| 88.80.150.82:8080 | ✓ 1528ms | ✓ 1900ms | 否 | 否 | ✓ 1834ms | https |
| 103.215.36.88:16455 | ✓ 1250ms | 否 | ✓ 1018ms | ✓ 1343ms | 否 | http |

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
