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

最后更新：2026-05-20 08:46:12 UTC（2026-05-20 16:46:12 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1876ms | ✓ 1498ms | 否 | ✓ 976ms | ✓ 755ms | http |
| 192.99.8.15:8850 | ✓ 1861ms | 否 | 否 | ✓ 1362ms | ✓ 1811ms | http |
| 202.28.194.139:31280 | ✓ 1445ms | 否 | ✓ 1332ms | ✓ 1742ms | ✓ 1755ms | http |
| 138.2.78.251:8100 | ✓ 1288ms | 否 | ✓ 1654ms | ✓ 1447ms | ✓ 1373ms | http |
| 138.2.92.70:8100 | ✓ 1532ms | 否 | 否 | ✓ 1357ms | ✓ 1052ms | http |
| 34.87.80.221:30000 | ✓ 765ms | ✓ 1754ms | ✓ 1295ms | ✓ 1113ms | ✓ 912ms | http |
| 176.111.37.5:39811 | ✓ 1071ms | 否 | ✓ 1086ms | 否 | ✓ 1495ms | http |
| 45.125.67.37:8443 | ✓ 1523ms | 否 | ✓ 1507ms | ✓ 1170ms | ✓ 1828ms | http |
| 113.160.132.26:8080 | ✓ 1304ms | ✓ 1772ms | ✓ 915ms | ✓ 1319ms | ✓ 1002ms | http |
| 176.111.37.216:39811 | ✓ 1037ms | 否 | ✓ 1380ms | 否 | ✓ 1799ms | http |
| 185.200.188.234:10001 | ✓ 1384ms | 否 | ✓ 1611ms | 否 | ✓ 1926ms | http |
| 78.186.118.164:3311 | ✓ 990ms | 否 | ✓ 1780ms | 否 | ✓ 1915ms | http |
| 150.107.140.238:3128 | ✓ 1426ms | 否 | ✓ 838ms | ✓ 1586ms | ✓ 1342ms | http |
| 152.67.191.232:6800 | ✓ 1402ms | 否 | ✓ 1033ms | ✓ 1416ms | ✓ 1202ms | http |
| 190.12.150.244:999 | ✓ 1690ms | 否 | ✓ 1136ms | ✓ 1797ms | ✓ 1863ms | http |
| 8.217.214.66:50000 | ✓ 1213ms | ✓ 1739ms | 否 | ✓ 1365ms | 否 | http |
| 106.10.55.212:1121 | 否 | ✓ 1158ms | ✓ 1868ms | ✓ 1722ms | ✓ 1042ms | http |
| 157.0.142.246:10057 | ✓ 1331ms | ✓ 1346ms | ✓ 1309ms | ✓ 1305ms | ✓ 1734ms | http |
| 38.19.41.227:999 | 否 | ✓ 1983ms | ✓ 1045ms | ✓ 1706ms | ✓ 1482ms | http |
| 74.208.192.81:3129 | ✓ 554ms | 否 | ✓ 536ms | ✓ 1440ms | ✓ 1294ms | http |
| 170.106.136.181:31002 | ✓ 276ms | ✓ 620ms | ✓ 529ms | ✓ 629ms | ✓ 445ms | http |
| 86.104.72.220:1081 | ✓ 949ms | ✓ 1778ms | ✓ 300ms | 否 | 否 | http |
| 152.42.170.187:9090 | ✓ 860ms | 否 | ✓ 1116ms | ✓ 1021ms | ✓ 826ms | http |
| 178.63.155.151:8898 | ✓ 1300ms | 否 | ✓ 1462ms | 否 | ✓ 1795ms | http |
| 148.230.4.241:999 | ✓ 925ms | ✓ 1374ms | ✓ 666ms | ✓ 1562ms | ✓ 1311ms | http |
| 122.2.48.121:8080 | ✓ 1191ms | 否 | ✓ 1141ms | ✓ 1202ms | ✓ 1220ms | http |
| 43.130.126.146:6688 | ✓ 559ms | 否 | ✓ 1223ms | 否 | ✓ 1435ms | http |
| 45.117.163.134:3128 | 否 | 否 | ✓ 1553ms | ✓ 1700ms | ✓ 1023ms | http |
| 8.217.78.60:8100 | 否 | 否 | ✓ 1013ms | ✓ 1999ms | ✓ 1121ms | http |
| 8.154.21.175:3128 | ✓ 1108ms | 否 | ✓ 879ms | 否 | ✓ 1420ms | http |
| 34.84.162.206:38080 | ✓ 1589ms | 否 | 否 | ✓ 1308ms | ✓ 937ms | http |
| 210.223.44.230:3128 | 否 | ✓ 996ms | ✓ 1627ms | 否 | ✓ 1429ms | http |
| 114.214.165.78:10810 | ✓ 1383ms | 否 | ✓ 1215ms | ✓ 1633ms | 否 | http |
| 69.164.251.114:8080 | ✓ 1342ms | 否 | ✓ 1303ms | 否 | ✓ 1736ms | http |
| 57.129.144.178:40000 | ✓ 878ms | ✓ 1913ms | ✓ 1816ms | 否 | ✓ 1583ms | http |
| 152.70.91.193:40000 | ✓ 1421ms | 否 | 否 | ✓ 1761ms | ✓ 1310ms | http |
| 88.248.130.191:3311 | ✓ 1161ms | 否 | ✓ 1161ms | 否 | ✓ 1702ms | http |
| 38.180.2.107:3128 | ✓ 1233ms | ✓ 1928ms | ✓ 1955ms | 否 | 否 | http |
| 209.97.149.157:80 | ✓ 362ms | ✓ 1462ms | ✓ 1275ms | ✓ 1408ms | ✓ 1372ms | http |
| 147.45.78.89:1080 | ✓ 1038ms | 否 | ✓ 899ms | 否 | ✓ 1452ms | http |
| 34.96.238.40:8080 | ✓ 1242ms | ✓ 1263ms | 否 | ✓ 1833ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1313ms | 否 | ✓ 1355ms | 否 | ✓ 1655ms | http |
| 103.247.21.220:4357 | 否 | 否 | ✓ 1375ms | ✓ 1527ms | ✓ 1472ms | http |
| 218.108.131.186:17890 | ✓ 961ms | ✓ 1077ms | ✓ 1242ms | ✓ 1139ms | ✓ 1542ms | http |
| 118.113.244.237:1080 | ✓ 1033ms | ✓ 1705ms | ✓ 1551ms | 否 | ✓ 1094ms | http |
| 223.16.170.103:80 | ✓ 1276ms | 否 | 否 | ✓ 993ms | ✓ 1058ms | http |
| 174.137.134.182:2999 | 否 | ✓ 1809ms | ✓ 975ms | 否 | ✓ 1148ms | http |
| 115.231.181.40:8128 | ✓ 1182ms | ✓ 1173ms | 否 | ✓ 1168ms | ✓ 952ms | http |
| 103.189.197.43:7778 | 否 | 否 | ✓ 1343ms | ✓ 1155ms | ✓ 1453ms | http |
| 121.130.177.28:8888 | ✓ 1453ms | 否 | ✓ 1485ms | ✓ 1628ms | 否 | http |
| 119.23.68.90:9003 | ✓ 888ms | ✓ 1188ms | ✓ 983ms | ✓ 1063ms | ✓ 861ms | http |
| 159.223.41.216:9090 | ✓ 700ms | 否 | ✓ 817ms | ✓ 1010ms | ✓ 880ms | http |
| 81.30.156.115:8080 | ✓ 985ms | 否 | ✓ 1652ms | 否 | ✓ 1721ms | http |
| 185.104.249.25:3128 | ✓ 1113ms | 否 | ✓ 1799ms | 否 | ✓ 1635ms | http |
| 3.101.133.120:80 | ✓ 812ms | ✓ 1923ms | ✓ 914ms | ✓ 1080ms | ✓ 957ms | http |
| 168.138.171.204:8100 | ✓ 1571ms | 否 | 否 | ✓ 1377ms | ✓ 877ms | http |
| 103.174.236.120:8092 | 否 | 否 | ✓ 1221ms | ✓ 1394ms | ✓ 1437ms | http |
| 64.188.77.26:3128 | 否 | ✓ 1960ms | ✓ 1039ms | ✓ 1896ms | 否 | http |
| 129.212.224.122:3128 | ✓ 693ms | 否 | ✓ 729ms | ✓ 988ms | ✓ 801ms | http |
| 64.188.77.221:3128 | 否 | 否 | ✓ 1257ms | ✓ 1889ms | ✓ 1513ms | http |
| 8.210.161.8:8100 | 否 | ✓ 1834ms | ✓ 1770ms | ✓ 1186ms | 否 | http |
| 114.214.163.108:6789 | ✓ 1317ms | ✓ 1340ms | 否 | ✓ 1460ms | ✓ 1187ms | http |
| 47.242.163.146:8100 | ✓ 1357ms | ✓ 1871ms | ✓ 1889ms | 否 | 否 | http |
| 192.81.129.252:3136 | 否 | 否 | ✓ 765ms | ✓ 617ms | ✓ 791ms | http |
| 8.218.174.172:8100 | 否 | ✓ 1306ms | ✓ 1721ms | 否 | ✓ 1327ms | http |
| 47.241.32.135:8100 | 否 | 否 | ✓ 1154ms | ✓ 1819ms | ✓ 1770ms | http |
| 185.21.15.206:3128 | ✓ 1508ms | 否 | ✓ 1488ms | 否 | ✓ 1832ms | http |
| 158.160.215.167:8124 | ✓ 1442ms | ✓ 1945ms | 否 | 否 | ✓ 1890ms | http |
| 144.31.25.69:21064 | ✓ 1419ms | 否 | ✓ 1279ms | ✓ 1885ms | 否 | http |
| 103.167.171.149:7778 | ✓ 1777ms | 否 | ✓ 1389ms | ✓ 1428ms | 否 | http |
| 27.254.99.183:8118 | ✓ 1457ms | 否 | ✓ 1068ms | 否 | ✓ 948ms | http |
| 103.210.160.62:7789 | 否 | ✓ 910ms | ✓ 1609ms | ✓ 1240ms | ✓ 1020ms | http |
| 152.32.132.190:7890 | ✓ 1251ms | ✓ 1833ms | 否 | 否 | ✓ 1032ms | http |
| 85.192.29.60:3128 | ✓ 1057ms | ✓ 1926ms | ✓ 1328ms | ✓ 1945ms | ✓ 1726ms | http |
| 138.68.101.169:3128 | ✓ 1264ms | 否 | ✓ 1979ms | ✓ 1865ms | 否 | http |
| 61.52.131.172:8443 | ✓ 926ms | ✓ 1116ms | ✓ 892ms | ✓ 1113ms | ✓ 948ms | http |
| 103.164.180.185:8060 | 否 | 否 | ✓ 1881ms | ✓ 1852ms | ✓ 1825ms | http |

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
