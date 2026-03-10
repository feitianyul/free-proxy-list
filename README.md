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

最后更新：2026-03-10 14:03:06 UTC（2026-03-10 22:03:06 UTC+8）

**代理总数：66**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 154.3.236.202:3128 | ✓ 437ms | 否 | ✓ 948ms | ✓ 1446ms | ✓ 833ms | http |
| 205.209.118.30:3138 | ✓ 371ms | 否 | ✓ 894ms | ✓ 1263ms | ✓ 1173ms | http |
| 45.136.131.47:8443 | ✓ 1919ms | 否 | ✓ 311ms | ✓ 858ms | ✓ 1198ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 879ms | ✓ 926ms | ✓ 713ms | http |
| 152.42.213.210:8080 | ✓ 1427ms | 否 | ✓ 1890ms | ✓ 1071ms | ✓ 858ms | http |
| 45.136.198.40:3128 | ✓ 824ms | 否 | ✓ 1926ms | 否 | ✓ 1690ms | http |
| 120.198.141.84:22222 | ✓ 1211ms | ✓ 1257ms | ✓ 929ms | 否 | ✓ 984ms | http |
| 120.238.159.229:22222 | ✓ 922ms | ✓ 1266ms | ✓ 1014ms | ✓ 1218ms | ✓ 953ms | http |
| 152.70.98.46:8888 | ✓ 1440ms | 否 | ✓ 1454ms | ✓ 1146ms | ✓ 778ms | http |
| 101.47.73.135:3128 | ✓ 1906ms | 否 | 否 | ✓ 1912ms | ✓ 1997ms | http |
| 138.124.53.25:7443 | 否 | ✓ 1541ms | ✓ 984ms | 否 | ✓ 999ms | http |
| 14.225.222.164:7890 | 否 | 否 | ✓ 1320ms | ✓ 1823ms | ✓ 1965ms | http |
| 162.240.154.26:3128 | ✓ 729ms | 否 | 否 | ✓ 1308ms | ✓ 1349ms | http |
| 115.231.181.40:8128 | ✓ 997ms | ✓ 1761ms | 否 | ✓ 1199ms | ✓ 1963ms | http |
| 39.104.201.40:7890 | ✓ 935ms | ✓ 1218ms | ✓ 1245ms | ✓ 1248ms | ✓ 963ms | http |
| 202.155.12.161:443 | ✓ 1865ms | ✓ 1536ms | ✓ 1057ms | ✓ 1153ms | ✓ 1590ms | http |
| 101.43.255.96:80 | ✓ 1287ms | 否 | ✓ 1619ms | 否 | ✓ 1701ms | http |
| 116.80.96.100:3172 | 否 | 否 | ✓ 1646ms | ✓ 1954ms | ✓ 1767ms | http |
| 35.225.22.61:80 | ✓ 602ms | 否 | ✓ 948ms | ✓ 1007ms | ✓ 1116ms | http |
| 81.70.169.194:80 | 否 | ✓ 1254ms | ✓ 1254ms | ✓ 1226ms | ✓ 1071ms | http |
| 91.107.141.42:8081 | ✓ 1677ms | 否 | ✓ 1172ms | 否 | ✓ 1510ms | http |
| 201.182.85.185:999 | ✓ 1549ms | ✓ 1931ms | 否 | ✓ 1826ms | 否 | http |
| 172.212.68.37:3128 | ✓ 355ms | 否 | ✓ 916ms | ✓ 1377ms | ✓ 1345ms | http |
| 150.107.140.238:3128 | ✓ 1917ms | 否 | 否 | ✓ 1882ms | ✓ 1189ms | http |
| 116.80.82.227:3172 | ✓ 1890ms | 否 | ✓ 1578ms | 否 | ✓ 1901ms | http |
| 116.80.82.224:3172 | ✓ 1889ms | 否 | ✓ 1579ms | 否 | ✓ 1692ms | http |
| 120.238.159.228:22222 | ✓ 929ms | 否 | 否 | ✓ 1288ms | ✓ 916ms | http |
| 120.238.159.230:22222 | ✓ 963ms | 否 | ✓ 1100ms | 否 | ✓ 933ms | http |
| 165.227.5.10:8888 | ✓ 648ms | ✓ 1750ms | ✓ 687ms | 否 | ✓ 1592ms | http |
| 209.38.51.97:3128 | ✓ 643ms | ✓ 1321ms | ✓ 1906ms | ✓ 1345ms | ✓ 962ms | http |
| 45.136.130.223:8443 | ✓ 977ms | 否 | ✓ 310ms | ✓ 1014ms | ✓ 805ms | http |
| 47.77.193.180:1080 | ✓ 1389ms | 否 | ✓ 1526ms | ✓ 1804ms | ✓ 1147ms | http |
| 113.59.32.163:22222 | ✓ 1403ms | ✓ 1947ms | 否 | 否 | ✓ 1355ms | http |
| 117.159.239.49:22222 | ✓ 1001ms | ✓ 1162ms | ✓ 901ms | ✓ 1111ms | ✓ 861ms | http |
| 120.240.29.173:22222 | 否 | ✓ 1502ms | ✓ 1041ms | ✓ 1228ms | ✓ 1046ms | http |
| 69.164.194.19:3128 | ✓ 541ms | ✓ 1807ms | ✓ 1502ms | 否 | ✓ 1475ms | http |
| 4.213.222.169:3128 | 否 | ✓ 1695ms | ✓ 1795ms | ✓ 1832ms | ✓ 1847ms | http |
| 120.238.159.234:22222 | ✓ 1002ms | ✓ 1285ms | ✓ 1009ms | ✓ 1151ms | ✓ 995ms | http |
| 95.3.9.78:8080 | ✓ 822ms | 否 | ✓ 1534ms | ✓ 1765ms | ✓ 1380ms | http |
| 168.235.110.63:3128 | ✓ 271ms | 否 | ✓ 1050ms | ✓ 1146ms | ✓ 899ms | http |
| 120.92.212.16:8890 | ✓ 957ms | ✓ 1230ms | 否 | 否 | ✓ 1011ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1485ms | ✓ 1250ms | ✓ 973ms | http |
| 183.249.5.214:22222 | ✓ 785ms | ✓ 1122ms | ✓ 919ms | ✓ 957ms | ✓ 738ms | http |
| 120.240.35.176:22222 | ✓ 946ms | ✓ 1354ms | ✓ 954ms | ✓ 1200ms | ✓ 974ms | http |
| 120.198.141.75:22222 | ✓ 1215ms | ✓ 1205ms | ✓ 1001ms | ✓ 1247ms | ✓ 976ms | http |
| 183.249.5.111:22222 | ✓ 1012ms | 否 | ✓ 998ms | 否 | ✓ 755ms | http |
| 120.240.35.177:22222 | ✓ 1070ms | ✓ 1226ms | ✓ 1103ms | 否 | ✓ 1665ms | http |
| 113.59.32.142:22222 | ✓ 1191ms | ✓ 1678ms | ✓ 1158ms | ✓ 1359ms | ✓ 1277ms | http |
| 222.184.48.235:22222 | ✓ 1184ms | 否 | 否 | ✓ 1303ms | ✓ 1322ms | http |
| 111.79.111.126:3128 | ✓ 1716ms | 否 | ✓ 1490ms | ✓ 1932ms | ✓ 1713ms | http |
| 103.39.51.190:8080 | ✓ 1859ms | 否 | ✓ 1876ms | ✓ 1879ms | ✓ 1484ms | http |
| 178.236.245.59:3128 | ✓ 803ms | 否 | ✓ 1539ms | 否 | ✓ 1838ms | http |
| 178.236.245.17:3128 | ✓ 1216ms | 否 | ✓ 1340ms | ✓ 1839ms | ✓ 1408ms | http |
| 121.138.61.193:8091 | 否 | 否 | ✓ 1221ms | ✓ 1160ms | ✓ 1086ms | http |
| 120.240.35.173:22222 | ✓ 1273ms | ✓ 1276ms | ✓ 1098ms | ✓ 1267ms | 否 | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 1875ms | ✓ 1553ms | ✓ 1023ms | http |
| 46.183.25.8:443 | ✓ 1198ms | 否 | ✓ 1671ms | 否 | ✓ 1548ms | http |
| 5.101.0.233:3128 | ✓ 827ms | 否 | ✓ 1259ms | 否 | ✓ 1721ms | http |
| 34.101.184.164:3128 | ✓ 1000ms | 否 | ✓ 1575ms | ✓ 1672ms | ✓ 1004ms | http |
| 61.52.131.172:8443 | ✓ 880ms | ✓ 1228ms | ✓ 911ms | ✓ 1199ms | ✓ 912ms | http |
| 117.159.239.58:22222 | ✓ 1053ms | ✓ 1193ms | ✓ 907ms | ✓ 1098ms | 否 | http |
| 109.234.38.35:3128 | ✓ 1551ms | 否 | ✓ 1526ms | ✓ 1393ms | ✓ 1167ms | http |
| 217.77.102.18:3128 | ✓ 1414ms | 否 | ✓ 1429ms | ✓ 1984ms | ✓ 1787ms | http |
| 113.177.131.2:3128 | ✓ 1970ms | 否 | ✓ 1346ms | ✓ 1558ms | ✓ 1095ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1071ms | ✓ 1350ms | ✓ 1091ms | 否 | http |
| 116.80.82.217:3172 | ✓ 1503ms | 否 | ✓ 1516ms | 否 | ✓ 1687ms | http |

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
