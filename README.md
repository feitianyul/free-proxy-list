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

最后更新：2026-03-06 11:32:58 UTC（2026-03-06 19:32:58 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 81 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 148ms | ✓ 1871ms | ✓ 916ms | ✓ 1057ms | ✓ 804ms | http |
| 152.42.195.165:8888 | ✓ 909ms | 否 | ✓ 920ms | ✓ 1336ms | ✓ 1015ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1438ms | ✓ 1142ms | ✓ 1203ms | ✓ 870ms | http |
| 159.223.42.219:3128 | ✓ 908ms | 否 | ✓ 1327ms | ✓ 1283ms | ✓ 1024ms | http |
| 178.236.245.17:3128 | ✓ 614ms | 否 | ✓ 1447ms | 否 | ✓ 1874ms | http |
| 178.236.245.59:3128 | ✓ 614ms | 否 | ✓ 1443ms | 否 | ✓ 1876ms | http |
| 154.37.208.132:30000 | ✓ 939ms | ✓ 1804ms | 否 | ✓ 1865ms | 否 | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 831ms | ✓ 1581ms | ✓ 828ms | http |
| 138.124.53.25:7443 | ✓ 1026ms | ✓ 1508ms | ✓ 1694ms | ✓ 1616ms | ✓ 1587ms | http |
| 14.56.177.44:3128 | ✓ 944ms | 否 | ✓ 1157ms | ✓ 1728ms | ✓ 1103ms | http |
| 120.92.212.16:7890 | ✓ 1228ms | 否 | ✓ 1848ms | ✓ 1823ms | ✓ 1118ms | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 1564ms | ✓ 1326ms | ✓ 989ms | http |
| 61.72.221.194:3128 | ✓ 1712ms | 否 | ✓ 1064ms | 否 | ✓ 1020ms | http |
| 107.174.80.186:3128 | ✓ 1954ms | 否 | ✓ 1385ms | ✓ 992ms | 否 | http |
| 192.166.82.55:1080 | ✓ 1856ms | 否 | ✓ 1878ms | ✓ 1772ms | 否 | http |
| 46.249.103.192:443 | ✓ 609ms | 否 | ✓ 938ms | ✓ 1910ms | 否 | http |
| 125.128.12.14:3128 | 否 | ✓ 1669ms | 否 | ✓ 1858ms | ✓ 1438ms | http |
| 91.107.175.112:10801 | ✓ 1366ms | 否 | ✓ 1157ms | 否 | ✓ 1621ms | http |
| 91.193.240.157:9877 | ✓ 807ms | 否 | ✓ 1010ms | 否 | ✓ 1438ms | http |
| 106.14.203.63:3333 | ✓ 1023ms | ✓ 1836ms | 否 | ✓ 1293ms | ✓ 1034ms | http |
| 81.70.169.194:80 | ✓ 1138ms | ✓ 1521ms | ✓ 1086ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 1584ms | ✓ 1475ms | ✓ 1118ms | ✓ 1488ms | ✓ 1192ms | http |
| 185.115.74.185:8080 | ✓ 700ms | ✓ 1613ms | ✓ 1604ms | 否 | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1092ms | ✓ 1294ms | ✓ 885ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1608ms | ✓ 1372ms | ✓ 1038ms | http |
| 47.77.193.180:1080 | ✓ 343ms | ✓ 981ms | ✓ 373ms | ✓ 871ms | ✓ 655ms | http |
| 125.128.12.144:3128 | ✓ 883ms | 否 | ✓ 928ms | ✓ 1322ms | ✓ 1301ms | http |
| 167.172.69.123:8080 | ✓ 978ms | 否 | ✓ 1246ms | ✓ 1306ms | ✓ 1179ms | http |
| 151.245.137.203:8085 | ✓ 976ms | 否 | ✓ 1216ms | 否 | ✓ 1780ms | http |
| 14.225.222.164:7890 | ✓ 1759ms | ✓ 1580ms | ✓ 1072ms | ✓ 1252ms | ✓ 1026ms | http |
| 14.225.217.30:7890 | ✓ 1848ms | 否 | 否 | ✓ 1721ms | ✓ 1423ms | http |
| 103.215.36.88:19160 | 否 | ✓ 1726ms | ✓ 1775ms | 否 | ✓ 1205ms | http |
| 116.80.82.231:3172 | ✓ 1644ms | 否 | ✓ 1643ms | 否 | ✓ 1803ms | http |
| 138.186.200.253:8081 | ✓ 801ms | ✓ 1277ms | ✓ 1035ms | 否 | 否 | http |
| 61.72.110.94:3128 | ✓ 1746ms | ✓ 1640ms | 否 | ✓ 1306ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1126ms | 否 | 否 | ✓ 1709ms | ✓ 1152ms | http |
| 89.185.85.138:1080 | ✓ 429ms | 否 | ✓ 1809ms | 否 | ✓ 1969ms | http |
| 61.72.110.54:3128 | ✓ 1522ms | 否 | ✓ 1401ms | 否 | ✓ 1642ms | http |
| 61.72.221.234:3128 | ✓ 1716ms | 否 | ✓ 1731ms | ✓ 1883ms | ✓ 1281ms | http |
| 14.56.107.244:3128 | ✓ 1144ms | ✓ 1998ms | ✓ 1831ms | 否 | 否 | http |
| 154.12.231.32:80 | ✓ 1676ms | 否 | ✓ 1507ms | ✓ 997ms | ✓ 998ms | http |
| 61.72.221.94:3128 | ✓ 1314ms | 否 | 否 | ✓ 1774ms | ✓ 1839ms | http |
| 104.129.203.244:10571 | ✓ 742ms | ✓ 981ms | ✓ 873ms | ✓ 1125ms | ✓ 902ms | http |
| 104.129.203.244:11465 | ✓ 743ms | 否 | ✓ 503ms | ✓ 938ms | ✓ 723ms | http |
| 42.115.72.27:2039 | ✓ 1729ms | 否 | 否 | ✓ 1978ms | ✓ 1738ms | http |
| 161.97.115.10:3128 | ✓ 880ms | 否 | ✓ 1314ms | 否 | ✓ 1202ms | http |
| 180.103.19.47:1080 | ✓ 1286ms | 否 | ✓ 1459ms | 否 | ✓ 1557ms | http |
| 42.115.72.27:2038 | ✓ 1701ms | 否 | ✓ 1967ms | ✓ 1967ms | ✓ 1741ms | http |
| 165.227.5.10:8888 | ✓ 777ms | 否 | ✓ 358ms | ✓ 963ms | ✓ 1037ms | http |
| 185.191.236.162:3128 | ✓ 1055ms | ✓ 1820ms | ✓ 1937ms | 否 | 否 | http |
| 104.129.203.244:11763 | ✓ 1540ms | ✓ 1188ms | ✓ 1809ms | ✓ 925ms | ✓ 797ms | http |
| 104.129.203.244:10785 | ✓ 1535ms | 否 | ✓ 953ms | ✓ 1003ms | ✓ 802ms | http |
| 172.212.68.37:3128 | ✓ 287ms | 否 | ✓ 975ms | ✓ 1771ms | ✓ 1062ms | http |
| 113.176.92.71:3128 | 否 | ✓ 1909ms | ✓ 1545ms | ✓ 1971ms | ✓ 1139ms | http |
| 154.64.240.39:1080 | ✓ 1558ms | ✓ 1999ms | ✓ 879ms | ✓ 1796ms | ✓ 1281ms | http |
| 89.110.116.166:8080 | ✓ 403ms | 否 | ✓ 1180ms | ✓ 1799ms | 否 | http |
| 45.136.198.40:3128 | ✓ 619ms | ✓ 1438ms | ✓ 1355ms | 否 | ✓ 1803ms | http |
| 222.127.55.155:8082 | ✓ 1979ms | 否 | ✓ 1509ms | ✓ 1986ms | 否 | http |
| 42.115.72.27:2049 | ✓ 1694ms | 否 | ✓ 1751ms | 否 | ✓ 1744ms | http |
| 103.139.138.194:3128 | ✓ 1259ms | 否 | ✓ 1693ms | ✓ 1718ms | 否 | http |
| 152.70.137.18:8888 | ✓ 893ms | ✓ 1592ms | 否 | ✓ 1684ms | 否 | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 1305ms | ✓ 1646ms | ✓ 1328ms | http |
| 146.190.232.76:3128 | ✓ 399ms | ✓ 1660ms | ✓ 455ms | ✓ 1243ms | ✓ 999ms | http |
| 201.144.25.226:3128 | ✓ 815ms | 否 | ✓ 1147ms | ✓ 1388ms | 否 | http |
| 187.216.141.46:3128 | ✓ 1445ms | 否 | ✓ 1933ms | ✓ 1951ms | ✓ 1677ms | http |
| 43.225.185.4:8000 | ✓ 1100ms | 否 | ✓ 1461ms | ✓ 1511ms | ✓ 1200ms | http |
| 193.124.190.224:53294 | ✓ 908ms | 否 | ✓ 1536ms | 否 | ✓ 1907ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1635ms | 否 | ✓ 1281ms | ✓ 953ms | http |
| 103.86.131.62:80 | ✓ 1237ms | 否 | 否 | ✓ 1478ms | ✓ 1381ms | http |
| 211.171.114.154:3128 | 否 | ✓ 1555ms | ✓ 1753ms | ✓ 1588ms | 否 | http |
| 194.59.204.87:9080 | ✓ 1516ms | 否 | ✓ 1800ms | 否 | ✓ 1970ms | http |
| 182.253.12.188:8080 | ✓ 1813ms | 否 | 否 | ✓ 1982ms | ✓ 1715ms | http |
| 104.129.203.244:11522 | ✓ 838ms | ✓ 1021ms | ✓ 1087ms | ✓ 998ms | ✓ 770ms | http |
| 104.129.203.244:11022 | ✓ 835ms | ✓ 1012ms | ✓ 949ms | ✓ 1032ms | ✓ 931ms | http |
| 193.108.118.190:8888 | ✓ 456ms | 否 | ✓ 654ms | 否 | ✓ 1440ms | http |
| 18.100.254.193:34241 | ✓ 1061ms | 否 | ✓ 1085ms | 否 | ✓ 1985ms | http |
| 167.172.69.123:80 | ✓ 1612ms | 否 | ✓ 1363ms | ✓ 1310ms | ✓ 1028ms | http |
| 77.221.18.126:3128 | ✓ 1252ms | 否 | ✓ 1688ms | 否 | ✓ 1760ms | http |
| 103.215.36.88:17565 | 否 | ✓ 1474ms | ✓ 1289ms | ✓ 1830ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1123ms | ✓ 1799ms | ✓ 1992ms | 否 | 否 | https |
| 85.9.195.140:1080 | ✓ 905ms | 否 | ✓ 1355ms | 否 | ✓ 1308ms | http |
| 59.153.16.105:20909 | ✓ 1904ms | 否 | 否 | ✓ 1865ms | ✓ 1799ms | http |

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
