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

最后更新：2026-03-08 15:37:18 UTC（2026-03-08 23:37:18 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 183.249.5.109:22222 | ✓ 830ms | 否 | ✓ 865ms | ✓ 1253ms | ✓ 906ms | http |
| 120.240.35.173:22222 | ✓ 1148ms | 否 | ✓ 1179ms | ✓ 1378ms | 否 | http |
| 113.59.32.162:22222 | 否 | ✓ 1565ms | ✓ 1237ms | 否 | ✓ 1093ms | http |
| 120.240.29.51:22222 | ✓ 1509ms | ✓ 1533ms | 否 | ✓ 1556ms | 否 | http |
| 205.209.118.30:3138 | ✓ 787ms | ✓ 1224ms | 否 | 否 | ✓ 939ms | http |
| 120.238.159.230:22222 | ✓ 1130ms | ✓ 1477ms | ✓ 1035ms | ✓ 1341ms | 否 | http |
| 120.238.159.228:22222 | ✓ 1115ms | ✓ 1381ms | ✓ 1012ms | ✓ 1333ms | ✓ 1103ms | http |
| 183.249.5.117:22222 | 否 | 否 | ✓ 846ms | ✓ 1164ms | ✓ 999ms | http |
| 120.232.242.119:22222 | ✓ 1096ms | ✓ 1354ms | 否 | ✓ 1292ms | 否 | http |
| 120.238.159.229:22222 | 否 | ✓ 1461ms | ✓ 1191ms | ✓ 1389ms | 否 | http |
| 46.183.25.8:443 | ✓ 1388ms | 否 | ✓ 1116ms | ✓ 1674ms | ✓ 1041ms | http |
| 172.212.68.37:3128 | ✓ 348ms | 否 | ✓ 783ms | ✓ 1395ms | ✓ 928ms | http |
| 152.42.213.210:80 | 否 | 否 | ✓ 1947ms | ✓ 1789ms | ✓ 1307ms | http |
| 101.43.255.96:80 | 否 | ✓ 1535ms | ✓ 1983ms | ✓ 1733ms | 否 | http |
| 35.225.22.61:80 | ✓ 328ms | ✓ 1166ms | 否 | 否 | ✓ 844ms | http |
| 168.235.110.63:3128 | ✓ 206ms | 否 | ✓ 1799ms | ✓ 1413ms | ✓ 756ms | http |
| 117.159.239.49:22222 | ✓ 1051ms | 否 | ✓ 1114ms | ✓ 1315ms | ✓ 997ms | http |
| 62.113.119.14:8080 | ✓ 961ms | 否 | ✓ 910ms | 否 | ✓ 1031ms | http |
| 81.70.169.194:80 | ✓ 1963ms | 否 | ✓ 1216ms | 否 | ✓ 1125ms | http |
| 202.155.12.161:443 | ✓ 1690ms | ✓ 1547ms | ✓ 1393ms | ✓ 1145ms | ✓ 1943ms | http |
| 120.240.35.175:22222 | ✓ 1064ms | ✓ 1553ms | ✓ 1068ms | ✓ 1309ms | ✓ 1087ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1424ms | 否 | ✓ 1496ms | ✓ 1133ms | http |
| 120.198.141.79:22222 | ✓ 1240ms | ✓ 1796ms | ✓ 1204ms | 否 | ✓ 1319ms | http |
| 120.240.35.178:22222 | ✓ 1194ms | ✓ 1468ms | ✓ 1097ms | 否 | ✓ 1090ms | http |
| 222.184.48.251:22222 | ✓ 1043ms | 否 | ✓ 1044ms | 否 | ✓ 1118ms | http |
| 194.213.18.200:443 | ✓ 507ms | ✓ 1605ms | ✓ 918ms | 否 | ✓ 1961ms | http |
| 178.236.245.59:3128 | ✓ 588ms | 否 | ✓ 913ms | ✓ 1781ms | ✓ 1451ms | http |
| 1.231.81.166:3128 | ✓ 1725ms | ✓ 1704ms | ✓ 1202ms | ✓ 1232ms | ✓ 1021ms | http |
| 162.248.165.72:1080 | ✓ 1547ms | 否 | ✓ 1163ms | 否 | ✓ 1925ms | http |
| 5.101.0.233:3128 | ✓ 1747ms | 否 | ✓ 1332ms | ✓ 1873ms | ✓ 1773ms | http |
| 209.38.51.97:3128 | 否 | ✓ 1089ms | ✓ 961ms | 否 | ✓ 965ms | http |
| 183.249.5.213:22222 | ✓ 855ms | ✓ 1301ms | ✓ 911ms | 否 | 否 | http |
| 117.159.239.51:22222 | ✓ 1059ms | ✓ 1251ms | 否 | 否 | ✓ 1022ms | http |
| 120.240.35.176:22222 | ✓ 1090ms | ✓ 1436ms | ✓ 1191ms | ✓ 1305ms | ✓ 1143ms | http |
| 113.59.32.141:22222 | ✓ 1198ms | ✓ 1555ms | ✓ 1077ms | ✓ 1365ms | ✓ 1106ms | http |
| 113.59.32.163:22222 | ✓ 1219ms | ✓ 1527ms | ✓ 1133ms | ✓ 1374ms | ✓ 1179ms | http |
| 121.230.8.245:1080 | ✓ 1269ms | ✓ 1981ms | ✓ 1957ms | 否 | 否 | http |
| 120.198.141.75:22222 | ✓ 1477ms | 否 | ✓ 1276ms | 否 | ✓ 1120ms | http |
| 147.75.202.36:10017 | ✓ 729ms | ✓ 1102ms | ✓ 575ms | ✓ 1220ms | 否 | http |
| 45.136.198.40:3128 | ✓ 760ms | ✓ 1881ms | ✓ 1511ms | 否 | ✓ 1553ms | http |
| 103.139.138.194:3128 | ✓ 1932ms | 否 | ✓ 1616ms | ✓ 1704ms | ✓ 1327ms | http |
| 121.128.121.54:3128 | ✓ 1861ms | ✓ 1352ms | ✓ 1949ms | ✓ 1246ms | ✓ 988ms | http |
| 20.210.76.178:8561 | ✓ 1460ms | ✓ 1132ms | ✓ 671ms | ✓ 990ms | ✓ 847ms | http |
| 20.27.15.49:8561 | ✓ 1460ms | ✓ 1145ms | ✓ 662ms | ✓ 992ms | ✓ 846ms | http |
| 20.210.76.175:8561 | ✓ 1463ms | ✓ 1364ms | ✓ 634ms | ✓ 996ms | ✓ 785ms | http |
| 20.210.76.104:8561 | ✓ 1460ms | ✓ 1408ms | ✓ 643ms | ✓ 988ms | ✓ 779ms | http |
| 125.128.12.14:3128 | ✓ 1498ms | ✓ 1471ms | ✓ 830ms | ✓ 1211ms | ✓ 1657ms | http |
| 103.215.36.88:17439 | 否 | ✓ 1546ms | ✓ 1256ms | ✓ 1592ms | ✓ 1162ms | http |
| 120.92.212.16:7890 | ✓ 1577ms | 否 | ✓ 1894ms | 否 | ✓ 1952ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1432ms | ✓ 1455ms | 否 | ✓ 1959ms | http |
| 152.42.213.210:8080 | ✓ 1707ms | 否 | ✓ 1827ms | ✓ 1431ms | ✓ 1506ms | http |
| 222.184.48.241:22222 | ✓ 1064ms | ✓ 1403ms | ✓ 1139ms | 否 | 否 | http |
| 183.249.5.214:22222 | ✓ 858ms | ✓ 1332ms | 否 | 否 | ✓ 898ms | http |
| 201.144.20.238:3128 | ✓ 641ms | 否 | ✓ 1476ms | ✓ 1403ms | ✓ 1126ms | http |
| 14.56.107.244:3128 | ✓ 1586ms | ✓ 1684ms | ✓ 933ms | ✓ 1277ms | ✓ 1151ms | http |
| 150.249.255.91:3128 | ✓ 1561ms | 否 | ✓ 1652ms | ✓ 1112ms | ✓ 1987ms | http |
| 103.39.51.190:8080 | ✓ 1983ms | 否 | ✓ 1474ms | ✓ 1951ms | ✓ 1789ms | http |
| 113.59.32.161:22222 | ✓ 1295ms | ✓ 1610ms | ✓ 1183ms | ✓ 1481ms | ✓ 1150ms | http |
| 61.72.221.94:3128 | ✓ 1988ms | 否 | ✓ 1389ms | ✓ 1425ms | ✓ 1012ms | http |
| 46.39.105.157:8080 | ✓ 513ms | 否 | ✓ 1016ms | ✓ 1381ms | ✓ 1042ms | http |
| 192.227.137.72:5050 | ✓ 476ms | ✓ 1188ms | ✓ 807ms | ✓ 1011ms | ✓ 944ms | http |
| 113.177.131.2:3128 | ✓ 1780ms | 否 | ✓ 1173ms | ✓ 1378ms | 否 | http |

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
